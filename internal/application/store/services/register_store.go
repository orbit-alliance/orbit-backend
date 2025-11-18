package services

import (
	"context"
	"fmt"

	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/store"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterStoreService struct {
	storeRepo  store.StoreRepository
	userRepo   user.UserRepository
	eventBus   *shared.EventBus
	ethGateway store.BlockchainGateway
}

func NewRegisterStoreService(
	storeRepo store.StoreRepository,
	userRepo user.UserRepository,
	eventBus *shared.EventBus,
	ethGateway store.BlockchainGateway,
) *RegisterStoreService {
	return &RegisterStoreService{
		storeRepo:  storeRepo,
		userRepo:   userRepo,
		eventBus:   eventBus,
		ethGateway: ethGateway,
	}
}

func (s *RegisterStoreService) RegisterStore(ctx context.Context, userId, walletAddress string) (*store.StoreInfoDTO, error) {
	//verify if store is already registered

	coinStatus, err := s.ethGateway.GetCoinsStatusByWallet(ctx, walletAddress)
	if err != nil {
		fmt.Println("Error getting coin status by wallet:", err)
		return nil, err
	}
	user, err := s.userRepo.FindByID(ctx, userId)
	if err != nil {
		fmt.Println("User to register store not found:", err)
		return nil, err
	}
	if user == nil {
		fmt.Println("User to register store not found")
		return nil, nil
	}
	newStore := store.NewStore(primitive.NewObjectID(), *user, walletAddress, *coinStatus)
	err = s.storeRepo.Save(ctx, newStore)
	if err != nil {
		fmt.Println("Error saving new store", err)
		return nil, err
	}
	storeDto := store.NewStoreInfoDTO(*newStore)
	s.eventBus.Publish(store.NewStoreRegistered(newStore))
	return storeDto, nil
}
