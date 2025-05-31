package coin

import "context"

type GoodActionRepository interface {
	Save(ctx context.Context, goodAction *GoodAction) error
	LoadAll(ctx context.Context) ([]GoodAction, error)
}
