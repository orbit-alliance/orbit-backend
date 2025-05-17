package user

// Deducts the transferred amount from the total balance, prioritizing the transferred balance and then the balance earned by shares
func (u *User) TransferCoins(to *User, amount uint64) (*TransferedCoins, error) {
	totalBalance := u.CoinStatus.EarnedByTransfer + u.CoinStatus.EarnedByActions
	if totalBalance < amount {
		return nil, ErrInsufficientBalance
	}

	u.CoinStatus.Transferred += amount

	transferUsed := min(u.CoinStatus.EarnedByTransfer, amount)
	u.CoinStatus.EarnedByTransfer -= transferUsed

	remaining := amount - transferUsed
	u.CoinStatus.EarnedByActions -= remaining

	return NewTransferedCoins(u, to, amount), nil
}

func (u *User) DoGoodAction(goodAction *UserGoodAction) *DidGoodAction {

	if goodAction == nil {
		return nil
	}
	if goodAction.UserID != u.ID {
		return nil
	}

	u.CoinStatus.EarnedByActions += uint64(goodAction.RewardAmount)
	return NewDidGoodAction(u, goodAction)
}

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}
