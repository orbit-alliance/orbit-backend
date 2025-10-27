package benefit

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SaleStatus uint8

const (
	DELIVERED SaleStatus = iota
	PENDING
	CANCELED
)

type BenefitSale struct {
	ID						primitive.ObjectID	`json:"id" bson:"_id"`
	UserId					string				`json:"user_id" bson:"user_id"`
	Username				string				`json:"username" bson:"username"`
	BenefitId				string				`json:"benefit_id" bson:"benefit_id"`
	BenefitName				string				`json:"benefit_name" bson:"benefit_name"`
	EarnedCoinsUsed			uint64				`json:"earned_coins_used" bson:"earned_coins_used"`
	TransferredCoinsUsed	uint64				`json:"transferred_coins_used" bson:"transferred_coins_used"`
	RequestedAt				time.Time			`json:"requested_at" bson:"requested_at"`
	UpdatedAt				time.Time			`json:"updated_at" bson:"updated_at"`
}

func NewBenefitSale(userId, username, benefitId, benefitName string, earnedCoinsUsed, transferredCoinsUsed uint64, requestedAt, updatedAt time.Time) *BenefitSale {
	return &BenefitSale{
		UserId:					userId,
		Username:				username,
		BenefitId:				benefitId,
		BenefitName:			benefitName,
		EarnedCoinsUsed:		earnedCoinsUsed,
		TransferredCoinsUsed:	transferredCoinsUsed,
		RequestedAt:			requestedAt,
		UpdatedAt:				updatedAt,
	}
}
