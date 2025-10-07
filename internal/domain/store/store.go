package store

import (
	"time"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
	"github.com/orbit-alliance/orbit-backend/internal/domain/coin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StoreStatus uint8

const (
	ACTIVE StoreStatus = iota
	INACTIVE
	SUSPENDED
)

// This benefit is just to start store implementation. It will be a domain
type Benefit struct {
	ID              primitive.ObjectID	`json:"id" bson:"_id"`
	Name			string				`json:"name" bson:"name"`
	Description		string				`json:"description" bson:"description"`
}

// Store will be a singleton in DB. It will have an id
type Store struct {
	ID                  primitive.ObjectID	`json:"id" bson:"_id"`
	AdmUser				user.User			`json:"adm_user" bson:"adm_user"`
	Benefits			[]Benefit			`json:"benefits" bson:"benefits"`
	BenefitQuantity		uint64				`json:"benefit_quantity" bson:"benefit_quantity"`
	Wallet				string				`json:"wallet" bson:"wallet"`
	CoinStatus			coin.CoinStatus		`json:"total_coin_amount" bson:"total_coin_amount"`
	LastUpdated			time.Time			`json:"last_updated" bson:"last_updated"`
	CreatedAt			time.Time			`json:"created_at" bson:"created_at"`
	TimeLastSale		time.Time			`json:"time_last_sale" bson:"time_last_sale"`
	Status				StoreStatus			`json:"store_status" bson:"store_status"`
}

// The idea is that this function would be called once.
func NewStore(id primitive.ObjectID, wallet string, coinStatus coin.CoinStatus) *Store {
	return &Store {
		ID: id,
		Wallet: wallet,
		CoinStatus: coinStatus,
	}
}
