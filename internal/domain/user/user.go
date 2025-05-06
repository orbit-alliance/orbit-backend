package user

import "time"

type User struct {
	ID42       string     `json:"id_42" bson:"id_42"`
	Wallet     string     `json:"wallet" bson:"wallet"`
	Username   string     `json:"username" bson:"username"`
	TokenStats TokenStats `json:"token_stats" bson:"token_stats"`
	NFTs       []NFT      `json:"nfts" bson:"nfts"`
	CreatedAt  time.Time  `json:"created_at" bson:"created_at"`
}

type TokenStats struct {
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
