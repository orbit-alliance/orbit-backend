package user

import (
	"context"

	"github.com/orbit-alliance/orbit-backend/internal/domain/coin"
)

type UserRepository interface {
	Save(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id string) (*User, error)
	FindByID42(ctx context.Context, id42 string) (*User, error)
	FindByWallet(ctx context.Context, wallet string) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	LoadAll(ctx context.Context) ([]*User, error)
}

type UserGoodActionRepository interface {
	Save(ctx context.Context, user *UserGoodAction) error
}

type TransferRepository interface {
	Save(ctx context.Context, transfer *coin.Transfer) error
}

type UserProjectRepository interface {
	Save(ctx context.Context, user *UserProject) error
	LoadByUserID(ctx context.Context, userID string) ([]*UserProject, error)
}

type GoodActionRepository interface {
	LoadByName(ctx context.Context, name string) (*coin.GoodAction, error)
}
