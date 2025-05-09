package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserBenefitPurchaseStatus uint8

const (
	PENDING UserBenefitPurchaseStatus = iota
	APPROVED
	DELIVERED
	REJECTED
)

type UserBenefitPurchase struct {
	ID                   primitive.ObjectID        `json:"id" bson:"id"`
	UserID               string                    `json:"user_id" bson:"user_id"`
	Username             string                    `json:"username" bson:"username"`
	BenefitID            primitive.ObjectID        `json:"benefit_id" bson:"benefit_id"`
	BenefitName          string                    `json:"benefit_name" bson:"benefit_name"`
	EarnedCoinsUsed      uint64                    `json:"earned_coins_used" bson:"earned_coins_used"`
	TransferredCoinsUsed uint64                    `json:"transferred_coins_used" bson:"transferred_coins_used"`
	PurchaseStatus       UserBenefitPurchaseStatus `json:"purchase_status" bson:"purchase_status"`
	RequestedAt          time.Time                 `json:"requested_at" bson:"requested_at"`
	UpdatedAt            time.Time                 `json:"updated_at" bson:"updated_at"`
}

func NewUserBenefitPurchase(userID, username string, benefitID primitive.ObjectID, benefitName string, earnedCoinsUsed, transferredCoinsUsed uint64) *UserBenefitPurchase {
	return &UserBenefitPurchase{
		ID:                   primitive.NewObjectID(),
		UserID:               userID,
		Username:             username,
		BenefitID:            benefitID,
		BenefitName:          benefitName,
		EarnedCoinsUsed:      earnedCoinsUsed,
		TransferredCoinsUsed: transferredCoinsUsed,
		PurchaseStatus:       PENDING,
		RequestedAt:          time.Now(),
	}
}
