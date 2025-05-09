package coin

import (
	"time"

	"github.com/orbit-alliance/orbit-backend/internal/domain/nft"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"` //Added
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
