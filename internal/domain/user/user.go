package user

import (
	"time"

	"github.com/orbit-alliance/orbit-backend/internal/domain/coin"
	"github.com/orbit-alliance/orbit-backend/internal/domain/nft"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status uint8

const (
	PENDING Status = iota
	APPROVED
	DELIVERED
	REJECTED
)

type User struct {
	ID         primitive.ObjectID `json:"id" bson:"id"`
	ID42       string             `json:"id_42" bson:"id_42"`
	Wallet     string             `json:"wallet" bson:"wallet"`
	Username   string             `json:"username" bson:"username"`
	CoinStatus coin.CoinStatus    `json:"coin_status" bson:"coin_status"`
	NFTs       []nft.NFT          `json:"nfts" bson:"nfts"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
}

type UserGoodAction struct {
	ID          primitive.ObjectID `json:"id" bson:"id"`
	UserId      string             `json:"user_id" bson:"user_id"`
	Username    string             `json:"username" bson:"username"`
	ActionId    string             `json:"action_id" bson:"action_id"`
	ActionName  string             `json:"action_name" bson:"action_name"`
	PerformedAt time.Time          `json:"performed_at" bson:"performed_at"`
}

type UserBenefitPurchase struct {
	ID                   primitive.ObjectID `json:"id" bson:"id"`
	UserID               string             `json:"user_id" bson:"user_id"`
	Username             string             `json:"username" bson:"username"`
	BenefitID            primitive.ObjectID `json:"benefit_id" bson:"benefit_id"`
	BenefitName          string             `json:"benefit_name" bson:"benefit_name"`
	EarnedCoinsUsed      uint64             `json:"earned_coins_used" bson:"earned_coins_used"`
	TransferredCoinsUsed uint64             `json:"transferred_coins_used" bson:"transferred_coins_used"`
	PurchaseStatus       Status             `json:"purchase_status" bson:"purchase_status"`
	RequestedAt          time.Time          `json:"requested_at" bson:"requested_at"`
	UpdatedAt            time.Time          `json:"updated_at" bson:"updated_at"`
}

func NewUser(id42, wallet, username string, coinStatus coin.CoinStatus, nfts []nft.NFT) *User {
	return &User{
		ID42:       id42,
		Wallet:     wallet,
		Username:   username,
		CoinStatus: coinStatus,
		NFTs:       nfts,
		CreatedAt:  time.Now(),
	}
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
