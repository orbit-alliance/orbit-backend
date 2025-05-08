package coin

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/orbit-alliance/orbit-backend/internal/domain/nft"
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

type GoodAction struct {
	ID              primitive.ObjectID `json:"id" bson:"id"`
	Name            string             `json:"name" bson:"name"`
	Description     string             `json:"description" bson:"description"`
	RewardAmount    uint64             `json:"reward_amount" bson:"reward_amount"`
	AllowedNFTTypes []nft.NFT          `json:"allowed_nft_types" bson:"allowed_nft_types"`
	BlockedNFTTypes []nft.NFT          `json:"blocked_nft_types" bson:"blocked_nft_types"`
}

type Benefit struct {
	ID              primitive.ObjectID `json:"id" bson:"id"`
	Name            string             `json:"name" bson:"name"`
	Description     string             `json:"description" bson:"description"`
	EarnedCost      uint64             `json:"earned_cost" bson:"earned_cost"`
	TransferredCost uint64             `json:"transferred_cost" bson:"transferred_cost"`
	TotalAvailable  uint64             `json:"total_available" bson:"total_available"`
	MaxPerUser      uint64             `json:"max_per_user" bson:"max_per_user"`
	AllowedNFTTypes []nft.NFT          `json:"allowed_nft_types" bson:"allowed_nft_types"`
	BlockedNFTTypes []nft.NFT          `json:"blocked_nft_types" bson:"blocked_nft_types"`
	CreatedAt       string             `json:"created_at" bson:"created_at"` //Added
}

type CoinStatus struct {
	ID               primitive.ObjectID `json:"id" bson:"id"` //Adicionado extra
	EarnedByActions  uint64             `json:"earned_by_actions" bson:"earned_by_actions"`
	EarnedByTransfer uint64             `json:"earned_by_transfer" bson:"earned_by_transfer"`
	Transferred      uint64             `json:"transferred" bson:"transferred"`
}

func NewTransfer(fromUserID, fromUsername, toUserID, toUsername string, amount uint64) *Transfer {
	return &Transfer{
		ID:           primitive.NewObjectID(),
		FromUserID:   fromUserID,
		FromUsername: fromUsername,
		ToUserID:     toUserID,
		ToUsername:   toUsername,
		Amount:       amount,
		PerformedAt:  time.Now(),
	}
}

func NewGoodAction(name, description string, rewardAmount uint64, allowedNFTTypes, blockedNFTTypes []nft.NFT) *GoodAction {
	return &GoodAction{
		ID:              primitive.NewObjectID(),
		Name:            name,
		Description:     description,
		RewardAmount:    rewardAmount,
		AllowedNFTTypes: allowedNFTTypes,
		BlockedNFTTypes: blockedNFTTypes,
	}
}

func NewBenefit(name, description string, earnedCost, transferredCost, totalAvailable, maxPerUser uint64, allowedNFTTypes, blockedNFTTypes []nft.NFT) *Benefit {
	return &Benefit{
		ID:              primitive.NewObjectID(),
		Name:            name,
		Description:     description,
		EarnedCost:      earnedCost,
		TransferredCost: transferredCost,
		TotalAvailable:  totalAvailable,
		MaxPerUser:      maxPerUser,
		AllowedNFTTypes: allowedNFTTypes,
		BlockedNFTTypes: blockedNFTTypes,
	}
}

func NewCoinStatus(earnedByActions, earnedByTransfer, transferred uint64) *CoinStatus {
	return &CoinStatus{
		ID:               primitive.NewObjectID(),
		EarnedByActions:  earnedByActions,
		EarnedByTransfer: earnedByTransfer,
		Transferred:      transferred,
	}
}
