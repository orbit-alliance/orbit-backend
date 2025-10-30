package store

import "context"

type StoreRepository interface {
	Save(ctx context.Context, store *Store) error
	FindByID(ctx context.Context, id string) (*Store, error)
	GetSingle(ctx context.Context) (*Store, error)
}
