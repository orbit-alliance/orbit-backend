package store

func (s *Store) RegisterStore(store *Store) (*StoreRegistered, error) {
	return NewStoreRegistered(store), nil
}
