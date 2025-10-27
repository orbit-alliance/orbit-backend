package store

import (
	"time"

	"github.com/orbit-alliance/orbit-backend/internal/domain/coin"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StoreStatus uint8

const (
	ACTIVE StoreStatus = iota
	INACTIVE
	SUSPENDED
)

type Store struct {
	ID              primitive.ObjectID	`json:"id" bson:"_id"`
	AdmUser         user.User			`json:"adm_user" bson:"adm_user"`
	BenefitQuantity uint64				`json:"benefit_quantity" bson:"benefit_quantity"`
	Wallet          string				`json:"wallet" bson:"wallet"`
	CoinStatus      coin.CoinStatus		`json:"total_coin_amount" bson:"total_coin_amount"`
	LastUpdated     time.Time			`json:"last_updated" bson:"last_updated"`
	CreatedAt       time.Time			`json:"created_at" bson:"created_at"`
	TimeLastSale    time.Time			`json:"time_last_sale" bson:"time_last_sale"`
	Status          StoreStatus			`json:"store_status" bson:"store_status"`
}

func NewStore(id primitive.ObjectID, admUser user.User, wallet string, coinStatus coin.CoinStatus) *Store {
	return &Store{
		ID:					id,
		AdmUser:			admUser,
		Wallet:				wallet,
		CoinStatus:			coinStatus,
		CreatedAt:			time.Now(),
		LastUpdated:		time.Now(),
		BenefitQuantity:	0,
		Status:				ACTIVE,
	}
}
