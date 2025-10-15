package services

import (
	"context"
	"fmt"
	"strconv"

	"github.com/orbit-alliance/orbit-backend/internal/domain/benefit"
	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/store"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterBenefitService struct {
	benefitRepo  benefit.BenefitRepository
	storeRepo  store.StoreRepository
	eventBus   *shared.EventBus
}

func NewRegisterBenefitService(
	benefitRepo benefit.BenefitRepository,
	storeRepo store.StoreRepository,
	eventBus *shared.EventBus,
) *RegisterBenefitService {
	return &RegisterBenefitService{
		benefitRepo: benefitRepo,
		storeRepo: storeRepo, 
		eventBus: eventBus,
	}
}

func (s *RegisterBenefitService) RegisterBenefit(ctx context.Context, storeId, name, description, imageUrl, category, totalAvailable, maxPeruser string) (*benefit.BenefitInfoDTO, error) {

	store, err := s.storeRepo.FindByID(ctx, storeId)

	if err != nil {
		fmt.Println("Store to register benefit not found:", err)
		return nil, err
	}

		
	totalAvailableInt, err := strconv.ParseUint(totalAvailable, 10, 64)
	if err != nil {
		fmt.Printf("Invalid total available number:", err);
		return nil, err
	}
	maxPerUserInt, err := strconv.ParseUint(maxPeruser, 10, 64) 
	if err != nil {
		fmt.Printf("Invalid max per user number:", err);
		return nil, err
	}

	newBenefit := benefit.NewBenefit(primitive.NewObjectID(), 
		name, description, imageUrl, category, store.AdmUser.Username,
		totalAvailableInt, maxPerUserInt)

	//STOPPED HERE


}
