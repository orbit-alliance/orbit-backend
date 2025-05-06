package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `json:"id" bson:"id"`
	ID42       string             `json:"id_42" bson:"id_42"`
	Wallet     string             `json:"wallet" bson:"wallet"`
	Username   string             `json:"username" bson:"username"`
	CoinStatus CoinStatus         `json:"coin_status" bson:"coin_status"`
	NFTs       []NFT              `json:"nfts" bson:"nfts"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
}

type CoinStatus struct {
	EarnedByActions  uint64 `json:"earned_by_actions" bson:"earned_by_actions"`
	EarnedByTransfer uint64 `json:"earned_by_transfer" bson:"earned_by_transfer"`
	Transferred      uint64 `json:"transferred" bson:"transferred"`
	Used             uint64 `json:"used" bson:"used"`
}

type NFT struct {
	ID        string      `json:"id" bson:"id"`
	Name      string      `json:"name" bson:"name"`
	CreatedAt time.Time   `json:"created_at" bson:"created_at"`
	Metadata  NFTMetadata `json:"metadata" bson:"metadata"`
}

type NFTMetadata struct {
	Title  string     `json:"title" bson:"title"`
	Image  string     `json:"image" bson:"image"`
	Expiry *time.Time `json:"expiry" bson:"expiry"`
}

func NewUser(id42, wallet, username string, coinStatus CoinStatus, nfts []NFT) *User {
	return &User{
		ID42:       id42,
		Wallet:     wallet,
		Username:   username,
		CoinStatus: coinStatus,
		NFTs:       nfts,
		CreatedAt:  time.Now(),
	}
}

func NewCoinStatus(earnedByActions, earnedByTransfer, transferred, used uint64) *CoinStatus {
	return &CoinStatus{
		EarnedByActions:  earnedByActions,
		EarnedByTransfer: earnedByTransfer,
		Transferred:      transferred,
		Used:             used,
	}
}

func NewNFT(id, name string, metadata NFTMetadata) *NFT {
	return &NFT{
		ID:        id,
		Name:      name,
		CreatedAt: time.Now(),
		Metadata:  metadata,
	}
}

func NewNFTMetadata(title, image string, expiry *time.Time) *NFTMetadata {
	return &NFTMetadata{
		Title:  title,
		Image:  image,
		Expiry: expiry,
	}
}
