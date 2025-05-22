package user

import (
	"context"
)

type TransferDTO struct {
	From   string
	To     string
	Amount uint64
}

type BlockchainGateway interface {
	PublishUserAction(ctx context.Context, payload *UserGoodAction) error
	TransferListener(ctx context.Context, handler func(event TransferDTO)) error
}

type BlockchainEventListener interface {
	Start(ctx context.Context) error
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
