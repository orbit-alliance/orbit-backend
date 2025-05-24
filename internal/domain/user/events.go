package user

import "github.com/orbit-alliance/orbit-backend/internal/domain/shared"

type User42Registered struct {
	shared.BaseEvent
	User *User
}

func (e User42Registered) EventType() string {
	return "user.User42Registered"
}
func NewUser42Registered(user *User) *User42Registered {
	return &User42Registered{
		BaseEvent: shared.NewBaseEvent(shared.NewMongoID()),
		User:      user,
	}
}

type UserWalletChanged struct {
	shared.BaseEvent
	user *User
}

func (e UserWalletChanged) EventType() string {
	return "user.UserWalletChanged"
}
func NewUserWalletChanged(user *User) *UserWalletChanged {
	return &UserWalletChanged{
		BaseEvent: shared.NewBaseEvent(shared.NewMongoID()),
		user:      user,
	}
}

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
		BaseEvent: shared.NewBaseEvent(shared.NewMongoID()),
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
		BaseEvent: shared.NewBaseEvent(shared.NewMongoID()),
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
		BaseEvent: shared.NewBaseEvent(shared.NewMongoID()),
		Action:    action,
	}
}
