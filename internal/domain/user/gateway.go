package user

import (
	"context"

	"github.com/orbit-alliance/orbit-backend/internal/domain/coin"
)

type TransferDTO struct {
	From   string
	To     string
	Amount uint64
}

type BlockchainGateway interface {
	PublishUserAction(ctx context.Context, payload *UserGoodAction) error
	TransferListener(ctx context.Context, handler func(event TransferDTO)) error
	GetCoinsStatusByWallet(ctx context.Context, wallet string) (*coin.CoinStatus, error)
}

type BlockchainEventListener interface {
	TransferListener(ctx context.Context) error
}

type UserProjectBonusDTO struct {
	ProjectName string
	Points      int
}

type UserLoggedDaysDTO struct {
	Date string
}

type UserPresenceInEventDTO struct {
	Name        string
	Date        string
	HasPresence bool
}

type Api42Gateway interface {
	GetRetroactiveBonusProject(userID string) ([]UserProjectBonusDTO, error)
	GetRetroactiveLoggedDays(userID string, startAt string) ([]UserLoggedDaysDTO, error)
	GetBasicUserInfo(ctx context.Context, token string) (ID42 string, login string, err error)
}

type ApiGoogleGateway interface {
	GetRetroativePresencesBy42ID(ID42 string) ([]UserPresenceInEventDTO, error)
}
