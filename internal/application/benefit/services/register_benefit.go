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
		fmt.Println("Invalid total available number:", err);
		return nil, err
	}
	maxPerUserInt, err := strconv.ParseUint(maxPeruser, 10, 64) 
	if err != nil {
		fmt.Println("Invalid max per user number:", err);
		return nil, err
	}
	newBenefit := benefit.NewBenefit(primitive.NewObjectID(), 
		name, description, imageUrl, category, store.AdmUser.Username,
		totalAvailableInt, maxPerUserInt)
	err = s.benefitRepo.Save(ctx, newBenefit)
	if (err != nil) {
		fmt.Println("Error saving benefit: ", err)
		return nil, err
	}
	benefitDto := benefit.NewBenefitInfoDTO(newBenefit.ID.Hex(), newBenefit.Name, 
		newBenefit.Description, newBenefit.ImageURL, newBenefit.Category,
		newBenefit.CreatedBy, newBenefit.CreatedAt.String(), newBenefit.LastUpdated.String(), 
		int64(newBenefit.EarnedCost), int64(newBenefit.TransferedCost), int64(newBenefit.TotalAvailable), int64(newBenefit.MaxPeruser),
		newBenefit.IsActive, newBenefit.Tags)
	s.eventBus.Publish(benefit.NewBenefitRegistered(newBenefit))
	return benefitDto, nil 
}
