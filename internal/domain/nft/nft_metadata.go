package nft

import (
	"time"
)

type NFTMetadata struct {
	Title  string     `json:"title" bson:"title"`
	Image  string     `json:"image" bson:"image"`
	Expiry *time.Time `json:"expiry" bson:"expiry"`
}

func NewNFTMetadata(title, image string, expiry *time.Time) NFTMetadata {
	return NFTMetadata{
		Title:  title,
		Image:  image,
		Expiry: expiry,
	}
}
