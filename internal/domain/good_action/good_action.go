package good_action

import (
	"github.com/orbit-alliance/orbit-backend/nft"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GoodAction struct {
	ID              primitive.ObjectID `json:"id" bson:"id"`
	Name            string             `json:"name" bson:"name"`
	Description     string             `json:"description" bson:"description"`
	RewardAmount    uint64             `json:"reward_amount" bson:"reward_amount"`
	AllowedNFTTypes []nft.NFT          `json:"allowed_nft_types" bson:"allowed_nft_types"`
	BlockedNFTTypes []nft.NFT          `json:"blocked_nft_types" bson:"blocked_nft_types"`
}
