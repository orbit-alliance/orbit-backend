package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserGoodAction struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	UserID       primitive.ObjectID `json:"user_id" bson:"user_id"`
	UserWallet   string             `json:"user_wallet" bson:"user_wallet"`
	Username     string             `json:"username" bson:"username"`
	ActionID     primitive.ObjectID `json:"action_id" bson:"action_id"`
	ActionName   string             `json:"action_name" bson:"action_name"`
	RewardAmount int64              `json:"reward_amount" bson:"reward_amount"`
	PerformedAt  time.Time          `json:"performed_at" bson:"performed_at"`
}

func NewUserGoodAction(
	userID primitive.ObjectID,
	userWallet string,
	username string,
	actionID primitive.ObjectID,
	actionName string,
	rewardAmount int64,
	performedAt time.Time,
) *UserGoodAction {
	return &UserGoodAction{
		ID:           primitive.NewObjectID(),
		UserID:       userID,
		UserWallet:   userWallet,
		Username:     username,
		ActionID:     actionID,
		ActionName:   actionName,
		RewardAmount: rewardAmount,
		PerformedAt:  performedAt,
	}
}
