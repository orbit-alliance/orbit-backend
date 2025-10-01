package store

import (
	"time"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
	"github.com/orbit-alliance/orbit-backend/internal/domain/coin"
)

type StoreStatus uint8

const (
	ACTIVE StoreStatus = iota
	INACTIVE
	SUSPENDED
)

type Store struct {
	AdmUser				user.User		`json:"adm_user" bson:"adm_user"`
	BenefitQuantity		uint64			`json:"benefit_quantity" bson:"benefit_quantity"`
	Wallet				string			`json:"wallet" bson:"wallet"`
	TotalCoinAmount		coin.CoinStatus	`json:"total_coin_amount" bson:"total_coin_amount"`
	LastUpdated			time.Time		`json:"last_updated" bson:"last_updated"`
	CreatedAt			time.Time		`json:"created_at" bson:"created_at"`
	TimeLastSale		time.Time		`json:"time_last_sale" bson:"time_last_sale"`
	Status				StoreStatus		`json:"store_status" bson:"store_status"`
}
