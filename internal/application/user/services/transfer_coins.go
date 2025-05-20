package services

import (
	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
)

type TransferCoinsService struct {
	userRepo user.UserRepository
	eventBus *shared.EventBus
}

func NewTransferCoinsService(userRepo user.UserRepository, eventBus *shared.EventBus) *TransferCoinsService {
	return &TransferCoinsService{
		userRepo: userRepo,
		eventBus: eventBus,
	}
}

// TO DO: TransferCoins transfers coins from one user to another
func (s *TransferCoinsService) TransferCoins(from, to string, amount uint64) error {
	return nil
}
