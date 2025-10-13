package services

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/store"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
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

	// store, err := s.storeRepo.FindByID(ctx, "1")

	// findUser
	// token42, err := s.api42Gateway.ExchangeCodeForToken(ctx, code42)
	// if err != nil {
	// 	fmt.Println("Error exchanging code for token:", err)
	// 	return nil, err
	// }

	// basicInfo, err := s.api42Gateway.GetBasicUserInfo(ctx, token42)
	// if err != nil {
	// 	fmt.Println("Error getting basic user info:", err)
	// 	return nil, err
	// }

	// fmt.Println(basicInfo)

	user, err := s.userRepo.FindByID(ctx, userId)

	if err != nil {
		fmt.Println("User to register store not found:", err)
		return nil, err
	}

	newStore := store.NewStore(primitive.NewObjectID(), *user, walletAddress, *coinStatus)

	s.storeRepo.Save(ctx, newStore)
	err = s.storeRepo.Save(ctx, newStore)
	if err != nil {
		fmt.Println("Error saving new store", err)
		return nil, err
	}

	storeDto := store.NewStoreInfoDTO(newStore.ID.Hex(), newStore.AdmUser.Username, newStore.CreatedAt.String(), newStore.LastUpdated.String(), int64(newStore.BenefitQuantity))

	s.eventBus.Publish(store.NewStoreRegistered(newStore))

	return storeDto, nil

	// STOPPED HERE
}
