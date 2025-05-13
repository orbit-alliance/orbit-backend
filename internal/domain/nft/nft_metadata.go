package nft

import (
	"time"
)

type NFTMetadata struct {
	Image  string     `json:"image" bson:"image"`
	Expiry *time.Time `json:"expiry" bson:"expiry"`
}

func NewNFTMetadata(image string, expiry *time.Time) *NFTMetadata {
	return &NFTMetadata{
		Image:  image,
		Expiry: expiry,
	}
}
