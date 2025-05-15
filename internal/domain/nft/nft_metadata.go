package nft

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NFTMetadata struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Image  string             `json:"image" bson:"image"`
	Expiry *time.Time         `json:"expiry" bson:"expiry"`
}

func NewNFTMetadata(image string, expiry *time.Time) *NFTMetadata {
	return &NFTMetadata{
		Image:  image,
		Expiry: expiry,
	}
}
