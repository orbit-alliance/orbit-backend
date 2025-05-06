package user

func (u *User) EarnCoins(amount uint64) {
	u.CoinStatus.EarnedByActions += amount
}

// Deducts the transferred amount from the total balance, prioritizing the transferred balance and then the balance earned by shares
func (u *User) TransferTokens(amount uint64) (*TransferedCoins, error) {
	totalBalance := u.CoinStatus.EarnedByTransfer + u.CoinStatus.EarnedByActions
	if totalBalance < amount {
		return nil, ErrInsufficientBalance
	}

	u.CoinStatus.Transferred += amount

	transferUsed := min(u.CoinStatus.EarnedByTransfer, amount)
	u.CoinStatus.EarnedByTransfer -= transferUsed

	remaining := amount - transferUsed
	u.CoinStatus.EarnedByActions -= remaining

	return NewTransferedCoins(u, u, amount), nil
}

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}
