package user

import (
	"time"

	"github.com/orbit-alliance/orbit-backend/internal/domain/coin"
	"github.com/orbit-alliance/orbit-backend/internal/domain/nft"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `json:"id" bson:"id"`
	ID42       string             `json:"id_42" bson:"id_42"`
	Wallet     string             `json:"wallet" bson:"wallet"`
	Username   string             `json:"username" bson:"username"`
	CoinStatus coin.CoinStatus    `json:"coin_status" bson:"coin_status"`
	NFTs       []nft.NFT          `json:"nfts" bson:"nfts"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
}

func NewUser(id42, wallet, username string, coinStatus coin.CoinStatus, nfts []nft.NFT) *User {
	return &User{
		ID42:       id42,
		Wallet:     wallet,
		Username:   username,
		CoinStatus: coinStatus,
		NFTs:       nfts,
		CreatedAt:  time.Now(),
	}
}
