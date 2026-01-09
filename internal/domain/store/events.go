package store

import (
	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
)

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

func (e ReceivedPurchaseCoins) EventType() string {
	return "store.ReceivedPurchaseCoins"
}
type ReceivedPurchaseCoins struct {
	shared.BaseEvent
	EarnedAmount		uint64
	TransferredAmount	uint64
	Buyer				*user.User
}

func NewReceivedPurchaseCoins(earnedAmount, transferredAmount uint64, buyer *user.User) *ReceivedPurchaseCoins {
	return &ReceivedPurchaseCoins {
		BaseEvent:			shared.NewBaseEvent(shared.NewMongoID()),
		EarnedAmount:		earnedAmount,
		TransferredAmount:	transferredAmount,
		Buyer:				buyer,
	}
}
