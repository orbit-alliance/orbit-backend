package store

// import (
// 	"time"
// )

func (s *Store) RegisterStore(store *Store) (*StoreRegistered, error) {
	return NewStoreRegistered(store), nil
}
