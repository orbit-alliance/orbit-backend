package coin

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CoinStatus struct {
	ID               primitive.ObjectID `json:"id" bson:"_id"`
	EarnedByActions  uint64             `json:"earned_by_actions" bson:"earned_by_actions"`
	EarnedByTransfer uint64             `json:"earned_by_transfer" bson:"earned_by_transfer"`
}

func NewCoinStatus(earnedByActions, earnedByTransfer uint64) *CoinStatus {
	return &CoinStatus{
		ID:               primitive.NewObjectID(),
		EarnedByActions:  earnedByActions,
		EarnedByTransfer: earnedByTransfer,
	}
}
