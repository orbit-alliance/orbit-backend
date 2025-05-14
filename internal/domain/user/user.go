package user

import (
	"time"

	"github.com/orbit-alliance/orbit-backend/internal/domain/coin"
	"github.com/orbit-alliance/orbit-backend/internal/domain/nft"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	ID42       string             `json:"id_42" bson:"_id_42"`
	Wallet     string             `json:"wallet" bson:"_wallet"`
	Username   string             `json:"username" bson:"_username"`
	CoinStatus coin.CoinStatus    `json:"coin_status" bson:"_coin_status"`
	NFTs       []nft.NFT          `json:"nfts" bson:"_nfts"`
	CreatedAt  time.Time          `json:"created_at" bson:"_created_at"`
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
