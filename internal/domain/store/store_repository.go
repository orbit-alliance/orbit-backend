package store

import "context"

type StoreRepository interface {
	Save(ctx context.Context, store *Store) error
	UpdateBenefit(ctx context.Context, store *Store, index int) error
	FindByID(ctx context.Context, id string) (*Store, error)
}
