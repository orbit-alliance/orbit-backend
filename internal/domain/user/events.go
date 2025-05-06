package user

import "github.com/orbit-alliance/orbit-backend/internal/domain/shared"

// TransferedCoins is the event that is triggered when a user transfers coins
type TransferedCoins struct {
	shared.BaseEvent
	Amount uint64
	From   *User
	To     *User
}

func (e TransferedCoins) EventType() string {
	return "user.TransferedCoins"
}

func NewTransferedCoins(from, to *User, amount uint64) *TransferedCoins {
	return &TransferedCoins{
		BaseEvent: shared.NewBaseEvent(from.ID),
		Amount:    amount,
		From:      from,
		To:        to,
	}
}
