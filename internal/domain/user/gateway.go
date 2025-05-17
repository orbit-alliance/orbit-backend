package user

import "context"

type BlockchainGateway interface {
	PublishUserAction(ctx context.Context, payload *UserGoodAction) error
}

type UserProjectBonusDTO struct {
	ProjectName string
	Points      int
}

type UserLoggedDaysDTO struct {
	Date string
}

type Api42Gateway interface {
	GetRetroactiveBonusProject(userID string) ([]UserProjectBonusDTO, error)
	GetRetroactiveLoggedDays(userID string, startAt string) ([]UserLoggedDaysDTO, error)
}
