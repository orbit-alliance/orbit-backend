package store

import "github.com/orbit-alliance/orbit-backend/internal/domain/shared"

type StoreRegistered struct {
	shared.BaseEvent
	Store *Store
}

func (e StoreRegistered) EventType() string {
	return "store.StoreRegistered"
}

func NewStoreRegistered(store *Store) *StoreRegistered {
	return &StoreRegistered {
		BaseEvent:	shared.NewBaseEvent(shared.NewMongoID()),
		Store:		store,
	}
}
