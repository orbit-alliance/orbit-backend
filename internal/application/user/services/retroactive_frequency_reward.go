package services

import (
	"context"
	"time"

	"github.com/orbit-alliance/orbit-backend/internal/domain/coin"
	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
)

type RetroactiveFrequencyRewardService struct {
	userRepo           user.UserRepository
	goodActionRepo     coin.GoodActionRepository
	userGoodActionRepo user.UserGoodActionRepository
	userGateway        user.Api42Gateway
	eventBus           *shared.EventBus
}

func NewRetroactiveFrequencyRewardService(
	userRepo user.UserRepository,
	goodActionRepo coin.GoodActionRepository,
	userGoodActionRepo user.UserGoodActionRepository,
	userGateway user.Api42Gateway,
	eventBus *shared.EventBus,
) *RetroactiveFrequencyRewardService {
	return &RetroactiveFrequencyRewardService{
		userRepo:           userRepo,
		goodActionRepo:     goodActionRepo,
		userGoodActionRepo: userGoodActionRepo,
		userGateway:        userGateway,
		eventBus:           eventBus,
	}
}

const (
	DAY_REWARD         = "Frequência Premiada"
	SEQUENCE_26_REWARD = "26 dias Consecutivos de Presença"
	SEQUENCE_42_REWARD = "42 dias Consecutivos de Presença"
)

// Não sei como deveria fazer para pegar as recompensas geradas do seed.
func getReward(rewardId string, rewardList []coin.GoodAction) coin.GoodAction {
	var rewardCopy coin.GoodAction

	for _, reward := range rewardList {
		if reward.Name == rewardId {
			rewardCopy = reward
			break
		}
	}
	return rewardCopy
}

func compareDates(start string, end string) bool {
	timeStart, err := time.Parse("2006-01-02", start)
	if err != nil {
		return false
	}
	timeEnd, err := time.Parse("2006-01-02", end)
	if err != nil {
		return false
	}
	timeStart = timeStart.AddDate(0, 0, 1)
	if timeStart.Equal(timeEnd) {
		return true
	}
	return false
}

func getRetroactiveReward(user user.User, logList []user.UserLoggedDaysDTO, rewardList []coin.GoodAction) (string, int, []coin.GoodAction) {
	var lastLogin = logList[0].Date
	var currentStreak int = user.CurrentStreak
	var lastDate string = user.LastLoginIn42
	var groupGoodActions []coin.GoodAction

	for _, log := range logList {
		if compareDates(lastDate, log.Date) {
			currentStreak++
		} else {
			currentStreak = 1
		}

		if currentStreak == 26 {
			groupGoodActions = append(groupGoodActions, getReward(SEQUENCE_26_REWARD, rewardList))
		} else if currentStreak == 42 {
			groupGoodActions = append(groupGoodActions, getReward(SEQUENCE_42_REWARD, rewardList))
		}
		groupGoodActions = append(groupGoodActions, getReward(DAY_REWARD, rewardList))
		lastDate = log.Date
	}
	return lastLogin, currentStreak, groupGoodActions
}

func (s *RetroactiveFrequencyRewardService) ApplyRetroactiveFrequencyRewardHandler(ctx context.Context, userID string) error {
	var usr, err = s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	logList, err := s.userGateway.GetRetroactiveLoggedDays(usr.ID42, usr.LastLoginIn42)
	if err != nil {
		return err
	}

	rewardList, err := s.goodActionRepo.LoadAll(ctx)
	if err != nil {
		return err
	}

	newLastLogin, newStreak, groupGoodActions := getRetroactiveReward(*usr, logList, rewardList)

	for _, action := range groupGoodActions {
		evt, userGoodAction, err := usr.DoFrequencyReward(newLastLogin, newStreak, action)
		if err != nil {
			return err
		}
		if evt != nil {
			s.eventBus.Publish(ctx, evt)

			if err := s.userRepo.Save(ctx, usr); err != nil {
				return err
			}
			if err := s.userGoodActionRepo.Save(ctx, userGoodAction); err != nil {
				return err
			}
		}
	}
	return nil
}
