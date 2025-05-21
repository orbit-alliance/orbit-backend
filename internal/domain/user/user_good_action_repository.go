package user

import "context"

type UserGoodActionRepository interface {
	Save(ctx context.Context, userGoodAction *UserGoodAction) error
}
