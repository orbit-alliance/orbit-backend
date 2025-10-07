package services

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/store"
)

type RegisterStoreService struct {
	storeRepo  store.StoreRepository
	eventBus   *shared.EventBus
	ethGateway store.BlockchainGateway
}

func NewRegisterStoreService(
	storeRepo store.StoreRepository,
	eventBus *shared.EventBus,
	ethGateway store.BlockchainGateway,
) *RegisterStoreService {
	return &RegisterStoreService{
		storeRepo:  storeRepo,
		eventBus:   eventBus,
		ethGateway: ethGateway,
	}
}

func (s *RegisterStoreService) RegisterStore(ctx context.Context, walletAddress string) (*store.StoreInfoDTO, error) {
	//verify if store is already registered

	coinStatus, err := s.ethGateway.GetCoinsStatusByWallet(ctx, walletAddress)
	if err != nil {
		fmt.Println("Error getting coin status by wallet:", err)
		return nil, err
	}

	// store, err := s.storeRepo.FindByID(ctx, "1")

	newStore := store.NewStore(primitive.NewObjectID(), walletAddress, *coinStatus)

	s.storeRepo.Save(ctx, newStore)
	err = s.storeRepo.Save(ctx, newStore)
	if err != nil {
		fmt.Println("Error saving new store", err)
		return nil, err
	}

	storeDto := store.NewStoreInfoDTO("01", "username", "created", "last", 0)

	s.eventBus.Publish(store.NewStoreRegistered(newStore))

	return storeDto, nil

	// STOPPED HERE
}
