package services

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/orbit-alliance/orbit-backend/internal/domain/coin"
	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
	"golang.org/x/sync/errgroup"
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

func getReward(rewardName string, rewardList []coin.GoodAction) (coin.GoodAction, error) {
	var rewardCopy coin.GoodAction
	var found bool
	for _, reward := range rewardList {
		if reward.Name == rewardName {
			rewardCopy = reward
			found = true
			break
		}
	}
	if !found {
		return coin.GoodAction{}, ErrGoodActionNotFound
	}
	return rewardCopy, nil
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
	return timeStart.Equal(timeEnd)
}

func getRetroactiveReward(user user.User, logList []user.UserLoggedDaysDTO, rewardList []coin.GoodAction) (string, int, []coin.GoodAction, error) {
	if len(logList) == 0 {
		return "", 0, nil, ErrUserNotHasStartDate
	}
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
			goodAction, err := getReward(SEQUENCE_26_REWARD, rewardList)
			if err != nil {
				return "", 0, nil, err
			}
			groupGoodActions = append(groupGoodActions, goodAction)
		} else if currentStreak == 42 {
			goodAction, err := getReward(SEQUENCE_42_REWARD, rewardList)
			if err != nil {
				return "", 0, nil, err
			}
			groupGoodActions = append(groupGoodActions, goodAction)
		}
		goodAction, err := getReward(DAY_REWARD, rewardList)
		if err != nil {
			return "", 0, nil, err
		}
		groupGoodActions = append(groupGoodActions, goodAction)
		lastDate = log.Date
	}
	return lastLogin, currentStreak, groupGoodActions, nil
}

func (s *RetroactiveFrequencyRewardService) retroactiveFrequencyRewardHandler(ctx context.Context, userID string, startAt string) error {
	var usr, err = s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	logList, err := s.userGateway.GetRetroactiveLoggedDays(usr.ID42, startAt)
	if err != nil {
		return err
	}

	rewardList, err := s.goodActionRepo.LoadAll(ctx)
	if err != nil {
		return err
	}

	newLastLogin, newStreak, groupGoodActions, err := getRetroactiveReward(*usr, logList, rewardList)
	if err != nil {
		return err
	}

	err = s.userRepo.SaveLogin(ctx, usr.ID, newLastLogin, newStreak)
	if err != nil {
		fmt.Printf("Error saving login for user %s: %v\n", usr.ID.Hex(), err)
		return err
	}

	totalTokens := uint64(0)
	for _, action := range groupGoodActions {
		evt, userGoodAction, err := usr.DoFrequencyReward(newLastLogin, newStreak, action)
		if err != nil {
			return err
		}
		if evt != nil {
			s.eventBus.Publish(evt)
			totalTokens += action.RewardAmount
			if err := s.userGoodActionRepo.Save(ctx, userGoodAction); err != nil {
				return err
			}
		}
	}

	if totalTokens > 0 {
		if err := s.userRepo.EarnTokens(ctx, usr.ID, totalTokens); err != nil {
			fmt.Printf("Error adding tokens for user %s: %v\n", usr.ID.Hex(), err)
			return err
		}
	}
	return nil
}

func (s *RetroactiveFrequencyRewardService) DailyFrequencyRewardJob() {
	ctx := context.Background()
	const maxConcurrency = 10

	// 1. Carrega todos os usuários
	users, err := s.userRepo.LoadAll(ctx)
	if err != nil {
		fmt.Printf("Erro ao carregar usuários: %v\n", err)
		return
	}

	// 2. Descobre a data‑alvo (“ontem” no fuso do servidor)
	startAt := time.Now().AddDate(0, 0, -1).Format("2006-01-02") // “YYYY-MM-DD”

	// 3. Cria errgroup com limite de concorrência
	g, gctx := errgroup.WithContext(ctx)
	sem := make(chan struct{}, maxConcurrency)

	for _, u := range users {
		u := u            // captura
		sem <- struct{}{} // bloqueia se já tiver 10 goroutines

		g.Go(func() error {
			defer func() { <-sem }() // libera vaga
			// IMPORTANTE: use o contexto do errgroup (propaga cancelamentos)
			if err := s.retroactiveFrequencyRewardHandler(gctx, u.ID.Hex(), startAt); err != nil {
				// loga mas deixa o errgroup propagar
				fmt.Printf("Erro em usuário %s: %v\n", u.ID.Hex(), err)
				return err
			}
			return nil
		})
	}

	// 4. Aguarda todas concluírem; se alguma falhou, imprime
	if err := g.Wait(); err != nil {
		fmt.Printf("DailyFrequencyRewardJob terminou com erro: %v\n", err)
	} else {
		fmt.Println("DailyFrequencyRewardJob concluído com sucesso.")
	}
}

func (s *RetroactiveFrequencyRewardService) ApplyRetroactiveFrequencyReward(ctx context.Context, event shared.DomainEvent) {

	evt, ok := event.(*user.User42Registered)
	if !ok {
		return
	}

	userID := shared.ObjectIDToString(evt.User.ID)
	startDate := os.Getenv("PROJECT_START_DATE")
	if startDate == "" {
		panic("PROJECT_START_DATE environment variable is not set")
	}

	if err := s.retroactiveFrequencyRewardHandler(ctx, userID, startDate); err != nil {
		fmt.Printf("Error applying retroactive frequency reward for user %s: %v\n", userID, err)
		return
	}

}
