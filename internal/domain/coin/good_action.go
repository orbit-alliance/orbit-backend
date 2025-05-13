package coin

import (
	"github.com/orbit-alliance/orbit-backend/internal/domain/nft"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GoodAction struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	Name            string             `json:"name" bson:"name"`
	Description     string             `json:"description" bson:"description"`
	RewardAmount    uint64             `json:"reward_amount" bson:"reward_amount"`
	GivenNFT        *nft.NFT           `json:"given_nft" bson:"given_nft"`
	AllowedNFTTypes []nft.NFT          `json:"allowed_nft_types" bson:"allowed_nft_types"`
	BlockedNFTTypes []nft.NFT          `json:"blocked_nft_types" bson:"blocked_nft_types"`
}

func NewGoodAction(name, description string, rewardAmount uint64, givenNFT *nft.NFT, allowedNFTTypes, blockedNFTTypes []nft.NFT) *GoodAction {
	return &GoodAction{
		ID:              primitive.NewObjectID(),
		Name:            name,
		Description:     description,
		RewardAmount:    rewardAmount,
		GivenNFT:        givenNFT,
		AllowedNFTTypes: allowedNFTTypes,
		BlockedNFTTypes: blockedNFTTypes,
	}
}
