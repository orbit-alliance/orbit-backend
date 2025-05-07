package nft

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NFTMetadata struct {
	Title  string     `json:"title" bson:"title"`
	Image  string     `json:"image" bson:"image"`
	Expiry *time.Time `json:"expiry" bson:"expiry"`
}

type NFT struct {
	ID        primitive.ObjectID `json:"id" bson:"id"` //Changed
	Name      string             `json:"name" bson:"name"`
	Metadata  NFTMetadata        `json:"metadata" bson:"metadata"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

func NewNFT(name string, metadata NFTMetadata) *NFT {
	return &NFT{
		ID:        primitive.NewObjectID(),
		Name:      name,
		Metadata:  metadata,
		CreatedAt: time.Now(),
	}
}

func NewNFTMetadata(title, image string, expiry *time.Time) NFTMetadata {
	return NFTMetadata{
		Title:  title,
		Image:  image,
		Expiry: expiry,
	}
}
