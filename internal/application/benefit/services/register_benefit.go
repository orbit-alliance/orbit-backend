package services

import (
	"context"
	"fmt"
	"slices"
	"strconv"
	"time"

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

func (s *RegisterBenefitService) RegisterBenefit(
	ctx context.Context, storeId, name, description, imageUrl, category, totalAvailable, maxPeruser string, isActive bool) (*benefit.BenefitInfoDTO, error) {

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
		totalAvailableInt, maxPerUserInt, isActive)

	store.Benefits = append(store.Benefits, *newBenefit)
	err = s.storeRepo.Save(ctx, store)

	err = s.benefitRepo.Save(ctx, newBenefit)

	if (err != nil) {
		fmt.Println("Error adding benefit to store: ", err)
		return nil, err
	}

	benefitDto := benefit.NewBenefitInfoDTO(newBenefit.ID.Hex(), newBenefit.Name, 
		newBenefit.Description, newBenefit.ImageURL, newBenefit.Category,
		newBenefit.CreatedBy, newBenefit.CreatedAt.String(), newBenefit.LastUpdated.String(), 
		uint64(newBenefit.EarnedCost), uint64(newBenefit.TransferedCost), uint64(newBenefit.TotalAvailable), uint64(newBenefit.MaxPerUser),
		newBenefit.IsActive, newBenefit.Tags)
	s.eventBus.Publish(benefit.NewBenefitRegistered(newBenefit))
	return benefitDto, nil 
}

func (s *RegisterBenefitService) UpdateBenefit(
	ctx context.Context, storeId, benefitId, name, description, imageUrl, category,
	totalAvailable, maxPerUser string, isActive bool) (*benefit.BenefitInfoDTO, error) {

	store, err := s.storeRepo.FindByID(ctx, storeId)
	
	if err != nil {
		fmt.Println("Store to update benefit not found:", err)
		return nil, err
	}

	index := slices.IndexFunc(store.Benefits, func(b benefit.Benefit) bool {
		return b.ID.Hex() == benefitId
	})

	store.Benefits[index].Name = name
	store.Benefits[index].Description = description
	store.Benefits[index].ImageURL = imageUrl
	totalAvailableNum, err := strconv.ParseUint(totalAvailable, 10, 64)
	if err != nil {
		fmt.Println("Cannot convert total available:", err)
		return nil, err
	}
	maxPerUserNum, err := strconv.ParseUint(maxPerUser, 10, 64)
	if err != nil {
		fmt.Println("Cannot convert max per user:", err)
		return nil, err
	}
	store.Benefits[index].TotalAvailable = totalAvailableNum
	store.Benefits[index].MaxPerUser = maxPerUserNum
	store.Benefits[index].LastUpdated = time.Now()
	store.Benefits[index].IsActive = isActive
	
	err = s.storeRepo.Save(ctx, store)
	if err != nil {
		fmt.Println("Could not update changes:", err)
		return nil, err
	}

	benefitDTO := benefit.NewBenefitDTO(&store.Benefits[index])
	return benefitDTO, nil
}

func (s *RegisterBenefitService) DeleteBenefit(ctx context.Context, storeId, benefitId string) error {
	
	store, err := s.storeRepo.FindByID(ctx, storeId)

	if err != nil {
		fmt.Println("Store to register benefit not found:", err)
		return err
	}

	store.Benefits = slices.DeleteFunc(store.Benefits, func(b benefit.Benefit) bool {
			return b.ID.Hex() == benefitId 
		})

	err = s.storeRepo.Save(ctx, store)
	if err != nil {
		fmt.Println("Could not save store changes:", err)
		return err
	}
	
	return nil
}
