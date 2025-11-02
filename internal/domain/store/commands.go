package store

import (
	"github.com/orbit-alliance/orbit-backend/internal/domain/benefit"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
)

func (s *Store) RegisterStore(store *Store) (*StoreRegistered, error) {
	return NewStoreRegistered(store), nil
}
func (s *Store) ReceivePurchaseCoins(from *user.User, b *benefit.Benefit) (*ReceivedPurchaseCoins, error) {
	s.AdmUser.CoinStatus.EarnedByActions += b.EarnedCost
	s.AdmUser.CoinStatus.EarnedByTransfer += b.TransferredCost

	return NewReceivedPurchaseCoins(b.EarnedCost, b.TransferredCost, from), nil
}
