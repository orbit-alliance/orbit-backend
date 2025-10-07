package store

import (
	"context"

	"github.com/orbit-alliance/orbit-backend/internal/domain/coin"
)

type StoreInfoDTO struct {
	ID              string `json:"_id"`
	AdmUsername     string `json:"adm_username"`
	CreatedAt       string `json:"created_at"`
	LastUpdated     string `json:"last_updated"`
	BenefitQuantity int64  `json:"benefit_quantity"`
}

func NewStoreInfoDTO(id, admUsername, createdAt, lastUpdated string, benefitQuatity int64) *StoreInfoDTO {
	return &StoreInfoDTO{
		ID:              id,
		AdmUsername:     admUsername,
		CreatedAt:       createdAt,
		LastUpdated:     lastUpdated,
		BenefitQuantity: benefitQuatity,
	}
}

type BlockchainGateway interface {
	GetCoinsStatusByWallet(ctx context.Context, wallet string) (*coin.CoinStatus, error)
}
