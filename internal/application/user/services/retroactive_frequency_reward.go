package services

import (
	"context"
	"time"

	"github.com/orbit-alliance/orbit-backend/internal/domain/coin"
	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
)

/* TODO list

1. Pegar o dia atual
2. Pegar o ultimo dia presente na tabela do usuário
3. Chamar a função que pega os dias logados entre essas datas
4. Adicionar ao usuário uma quantidade de moedas igual a quantidade de dias logados
5. Atualizar o último dia logado do usuário

1. Pegar a sequencia atual de dias logados do usuário
2. Verificar quantos dias logados estão em sequência + somado a sequencia atual
3. Adicionar ao usuário a cada marcador "x" um bonus de "n" moedas
4. Atualizar a sequencia atual do usuário

*/

type RetroactiveFrequencyRewardService struct {
	userRepo       user.UserRepository
	goodActionRepo coin.GoodActionRepository
	userData       user.User
	userGateway    user.Gateway
	eventBus       *shared.EventBus
}

func NewRetroactiveFrequencyRewardService(
	userRepo user.UserRepository,
	goodActionRepo coin.GoodActionRepository,
	userData user.User,
	userGateway user.Gateway,
	eventBus *shared.EventBus,
) *RetroactiveFrequencyRewardService {
	return &RetroactiveFrequencyRewardService{
		userRepo:       userRepo,
		goodActionRepo: goodActionRepo,
		userData:       userData,
		userGateway:    userGateway,
		eventBus:       eventBus,
	}
}

const (
	DAY_REWARD         = "test0"
	SEQUENCE_26_REWARD = "test1"
	SEQUENCE_42_REWARD = "test2"
)

// Não sei como deveria fazer para pegar as recompensas geradas do seed.
func createReward(rewardId string, rewardList []coin.GoodAction) coin.GoodAction {
	var rewardCopy coin.GoodAction

	for _, reward := range rewardList {
		if reward.ID == rewardId {
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
			groupGoodActions = append(groupGoodActions, createReward(SEQUENCE_26_REWARD, rewardList))
		} else if currentStreak == 42 {
			groupGoodActions = append(groupGoodActions, createReward(SEQUENCE_42_REWARD, rewardList))
		}
		groupGoodActions = append(groupGoodActions, createReward(DAY_REWARD, rewardList))
		lastDate = log.Date
	}
	return lastLogin, currentStreak, groupGoodActions
}

func (s *RetroactiveFrequencyRewardService) ApplyRetroactiveFrequencyRewardHandler(ctx context.Context) error {
	var user user.User = s.userData

	logList, err := s.userGateway.GetRetroactiveLoggedDays(user.ID42, user.LastLoginIn42)
	if err != nil {
		return err
	}

	rewardList, err := s.goodActionRepo.LoadAll(ctx)
	if err != nil {
		return err
	}

	newLastLogin, newStreak, groupGoodActions := getRetroactiveReward(s.userData, logList, rewardList)
	user.LastLoginIn42 = newLastLogin
	user.CurrentStreak = newStreak
	// Descobrir como adicionar boas ações aos usuários.
	err = s.userRepo.Save(ctx, &user)
	if err != nil {
		return err
	}
	//s.eventBus.Subscribe("ctx", this)
	// Descobrir como faço para adicionar o evento.
	return nil
}
