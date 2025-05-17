package nft

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NFT struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Metadata    NFTMetadata        `json:"metadata" bson:"metadata"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
}

func NewNFT(name, description string, metadata NFTMetadata) *NFT {
	return &NFT{
		ID:        primitive.NewObjectID(),
		Name:      name,
		Metadata:  metadata,
		CreatedAt: time.Now(),
	}
}
