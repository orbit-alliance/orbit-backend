package user

import (
	"time"

	"github.com/orbit-alliance/orbit-backend/internal/domain/coin"
	"github.com/orbit-alliance/orbit-backend/internal/domain/nft"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                   primitive.ObjectID `json:"id" bson:"_id"`
	ID42                 string             `json:"id_42" bson:"id_42"`
	Wallet               string             `json:"wallet" bson:"wallet"`
	Username             string             `json:"username" bson:"username"`
	CoinStatus           coin.CoinStatus    `json:"coin_status" bson:"coin_status"`
	NFTs                 []nft.NFT          `json:"nfts" bson:"nfts"`
	LastLoginIn42        string             `json:"last_login_in_42" bson:"last_login_in_42"`
	CurrentStreak        int                `json:"current_streak" bson:"current_streak"`
	RewardedDays         int                `json:"rewarded_days" bson:"rewarded_days"`
	RewardedDaysRequired int                `json:"rewarded_days_required" bson:"rewarded_days_required"`
	CreatedAt            time.Time          `json:"created_at" bson:"created_at"`
}

func NewUser(id primitive.ObjectID, id42, wallet, username string, coinStatus coin.CoinStatus, nfts []nft.NFT) *User {
	return &User{
		ID:         id,
		ID42:       id42,
		Wallet:     wallet,
		Username:   username,
		CoinStatus: coinStatus,
		NFTs:       nfts,
		CreatedAt:  time.Now(),
	}
}
