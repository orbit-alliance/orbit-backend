package user

import "github.com/orbit-alliance/orbit-backend/internal/domain/shared"

type SendedCoins struct {
	shared.BaseEvent
	Amount uint64
	From   *User
	To     *User
}

func (e SendedCoins) EventType() string {
	return "user.SendedCoins"
}

func NewSendedCoins(from, to *User, amount uint64) *SendedCoins {
	return &SendedCoins{
		BaseEvent: shared.NewBaseEvent(from.ID),
		Amount:    amount,
		From:      from,
		To:        to,
	}
}

type ReceivedCoins struct {
	shared.BaseEvent
	Amount uint64
	From   *User
	To     *User
}

func (e ReceivedCoins) EventType() string {
	return "user.ReceivedCoins"
}
func NewReceivedCoins(from, to *User, amount uint64) *ReceivedCoins {
	return &ReceivedCoins{
		BaseEvent: shared.NewBaseEvent(to.ID),
		Amount:    amount,
		From:      from,
		To:        to,
	}
}

type DidGoodAction struct {
	shared.BaseEvent
	Action *UserGoodAction
}

func (e DidGoodAction) EventType() string {
	return "user.DidGoodAction"
}
func NewDidGoodAction(user *User, action *UserGoodAction) *DidGoodAction {
	return &DidGoodAction{
		BaseEvent: shared.NewBaseEvent(user.ID),
		Action:    action,
	}
}
