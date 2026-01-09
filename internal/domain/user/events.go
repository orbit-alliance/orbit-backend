package user

import (
	"time"

	"github.com/orbit-alliance/orbit-backend/internal/domain/benefit"
	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
)

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
	User *User
}

func (e UserWalletChanged) EventType() string {
	return "user.UserWalletChanged"
}
func NewUserWalletChanged(user *User) *UserWalletChanged {
	return &UserWalletChanged{
		BaseEvent: shared.NewBaseEvent(shared.NewMongoID()),
		User:      user,
	}
}

type SentCoins struct {
	shared.BaseEvent
	Amount uint64
	From   *User
	To     *User
}

func (e SentCoins) EventType() string {
	return "user.SentCoins"
}

func NewSentCoins(from, to *User, amount uint64) *SentCoins {
	return &SentCoins{
		BaseEvent: shared.NewBaseEvent(shared.NewMongoID()),
		Amount:    amount,
		From:      from,
		To:        to,
	}
}

type PurchasedUserBenefit struct {
	shared.BaseEvent
	User					*User
	Benefit					*benefit.Benefit
	EarnedCoinsUsed			uint64                  
	TransferredCoinsUsed	uint64                  
	// PurchaseStatus			UserBenefitPurchaseStatus // TODO check if purchase_status is needed
	RequestedAt				time.Time
	UpdatedAt				time.Time
}
func (e PurchasedUserBenefit) EventType() string {
	return "user.PurchasedUserBenefit"
}

func NewPurchasedUserBenefit(user *User, benefit *benefit.Benefit, earnedUsed, transfUsed uint64) *PurchasedUserBenefit {
	return &PurchasedUserBenefit{
		User: user,
		Benefit: benefit,
		EarnedCoinsUsed: earnedUsed,
		TransferredCoinsUsed: transfUsed,
		// PurchaseStatus: p.PurchaseStatus, // TODO check if purchase_status is needed
		RequestedAt: time.Now(),
		UpdatedAt: time.Now(),
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
