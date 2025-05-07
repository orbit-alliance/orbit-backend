package transfer

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transfer struct {
	ID           primitive.ObjectID `json:"id" bson:"id"`
	FromUserID   string             `json:"from_user_id" bson:"from_user_id"`
	FromUsername string             `json:"from_username" bson:"from_username"`
	ToUserID     string             `json:"to_user_id" bson:"to_user_id"`
	ToUsername   string             `json:"to_username" bson:"to_username"`
	Amount       uint64             `json:"amount" bson:"amount"`
	PerformedAt  time.Time          `json:"created_at" bson:"created_at"`
}
