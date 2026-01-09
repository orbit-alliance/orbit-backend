package store

import (
	"context"

	"github.com/orbit-alliance/orbit-backend/internal/domain/coin"
)

type StoreInfoDTO struct {
	ID              string		`json:"_id"`
	AdmUsername     string		`json:"adm_username"`
	CreatedAt       string		`json:"created_at"`
	LastUpdated     string		`json:"last_updated"`
	TimeLastSale	string		`json:"time_last_sale"`
	BenefitQuantity uint64		`json:"benefit_quantity"`
	Status			StoreStatus `json:"store_status"`
}

func NewStoreInfoDTO(store Store) *StoreInfoDTO {
	return &StoreInfoDTO{
		ID:					store.ID.Hex(),
		AdmUsername:		store.AdmUser.Username,
		CreatedAt:			store.CreatedAt.String(),
		LastUpdated:		store.LastUpdated.String(),
		TimeLastSale:		store.TimeLastSale.String(),
		BenefitQuantity:	store.BenefitQuantity,
		Status:				store.Status,
	}
}

type BlockchainGateway interface {
	GetCoinsStatusByWallet(ctx context.Context, wallet string) (*coin.CoinStatus, error)
}
