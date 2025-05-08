package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserGoodAction struct {
	ID          primitive.ObjectID `json:"id" bson:"id"`
	UserId      string             `json:"user_id" bson:"user_id"`
	Username    string             `json:"username" bson:"username"`
	ActionId    string             `json:"action_id" bson:"action_id"`
	ActionName  string             `json:"action_name" bson:"action_name"`
	PerformedAt time.Time          `json:"performed_at" bson:"performed_at"`
}

func NewUserGoodAction(userId, username, actionId, actionName string) *UserGoodAction {
	return &UserGoodAction{
		ID:          primitive.NewObjectID(),
		UserId:      userId,
		Username:    username,
		ActionId:    actionId,
		ActionName:  actionName,
		PerformedAt: time.Now(),
	}
}
