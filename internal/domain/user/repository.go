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
}

type UserGoodActionRepository interface {
	Save(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id string) (*coin.GoodAction, error)
	LoadAll(ctx context.Context) ([]coin.GoodAction, error)
}
