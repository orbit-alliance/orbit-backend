package user

func (u *User) EarnCoins(amount uint64) {
	u.CoinStatus.EarnedByActions += amount
}

// Deduz o valor transferido do saldo total, priorizando o saldo transferido e depois o saldo ganho por ações
func (u *User) TransferTokens(amount uint64) error {
	totalBalance := u.CoinStatus.EarnedByTransfer + u.CoinStatus.EarnedByActions
	if totalBalance < amount {
		return ErrInsufficientBalance
	}

	u.CoinStatus.Transferred += amount

	transferUsed := min(u.CoinStatus.EarnedByTransfer, amount)
	u.CoinStatus.EarnedByTransfer -= transferUsed

	remaining := amount - transferUsed
	u.CoinStatus.EarnedByActions -= remaining

	return nil
}

func (u *User) ReceiveTokens(amount uint64) {
	u.CoinStatus.EarnedByTransfer += amount
}

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}
