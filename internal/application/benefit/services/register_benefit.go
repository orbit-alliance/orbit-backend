package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/orbit-alliance/orbit-backend/internal/domain/benefit"
	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/store"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterBenefitService struct {
	benefitRepo  benefit.BenefitRepository
	storeRepo  store.StoreRepository
	userRepo	user.UserRepository
	eventBus   *shared.EventBus
}

func NewRegisterBenefitService(
	benefitRepo	benefit.BenefitRepository,
	storeRepo	store.StoreRepository,
	userRepo	user.UserRepository,
	eventBus	*shared.EventBus,
) *RegisterBenefitService {
	return &RegisterBenefitService{
		benefitRepo:	benefitRepo,
		storeRepo:		storeRepo,
		userRepo:		userRepo,
		eventBus:		eventBus,
	}
}

func isValidAdmUser(s *RegisterBenefitService, ctx context.Context, store store.Store, admUserId string) bool {
	user, err := s.userRepo.FindByID(ctx, admUserId)
	if err != nil {
		fmt.Println("Store adm user not found:", err)
		return false
	}
	if (store.AdmUser.ID != user.ID) {
		fmt.Println("User is not a Store adm")
		return false
	}
	return true
}

func (s *RegisterBenefitService) RegisterBenefit(
	ctx context.Context, 
	storeId, admUserId, name, description, imageUrl, category string, 
	earnedCost, transferedCost, totalAvailable, maxPerUser uint64, 
	status benefit.BenefitStatus, 
	tags []string) (*benefit.BenefitInfoDTO, error) {
	
	store, err := s.storeRepo.FindByID(ctx, storeId)
	if err != nil {
		fmt.Println("Store to register benefit not found:", err)
		return nil, err
	}
	if (!isValidAdmUser(s, ctx, *store, admUserId)) {
		return nil, errors.New("Could not validate user")
	}

	newBenefit := benefit.NewBenefit(primitive.NewObjectID(), 
		name, description, imageUrl, category, store.AdmUser.Username,
		earnedCost, transferedCost, totalAvailable, maxPerUser, status, tags)
	err = s.benefitRepo.Save(ctx, newBenefit)
	if (err != nil) {
		fmt.Println("Error trying to save benefit", err)
		return nil, err
	}

	benefitDto := benefit.NewBenefitInfoDTO(*newBenefit)
	s.eventBus.Publish(benefit.NewBenefitRegistered(newBenefit))
	return benefitDto, nil 
}

func (s *RegisterBenefitService) UpdateBenefit(
	ctx context.Context, storeId, admUserId, benefitId, name, description, imageUrl, category string, 
	totalAvailable, maxPerUser uint64, status benefit.BenefitStatus, tags []string) (*benefit.BenefitInfoDTO, error) {

	store, err := s.storeRepo.FindByID(ctx, storeId)
	if err != nil {
		fmt.Println("Store to register benefit not found:", err)
		return nil, err
	}
	if (!isValidAdmUser(s, ctx, *store, admUserId)) {
		return nil, errors.New("Could not validate user")
	}

	updateBenefit, err := s.benefitRepo.FindByID(ctx, benefitId)
	updateBenefit.Name = name
	updateBenefit.Description = description
	updateBenefit.ImageURL = imageUrl
	updateBenefit.Category = category
	updateBenefit.TotalAvailable = totalAvailable
	updateBenefit.MaxPerUser = maxPerUser
	updateBenefit.LastUpdated = time.Now()
	updateBenefit.Status = status
	updateBenefit.Tags = tags

	err = s.benefitRepo.Update(ctx, updateBenefit)
	if err != nil {
		fmt.Println("Could not update changes:", err)
		return nil, err
	}
	benefitDTO := benefit.NewBenefitInfoDTO(*updateBenefit)

	return benefitDTO, nil
}

func (s *RegisterBenefitService) DeleteBenefit(ctx context.Context, storeId, admUserId, benefitId string) (*benefit.BenefitInfoDTO, error) {
	
	store, err := s.storeRepo.FindByID(ctx, storeId)
	if err != nil {
		fmt.Println("Store to register benefit not found:", err)
		return nil, err
	}
	if (!isValidAdmUser(s, ctx, *store, admUserId)) {
		return nil, errors.New("Could not validate user")
	}

	deleteBenefit, err := s.benefitRepo.FindByID(ctx, benefitId)
	if err != nil {
		fmt.Println("Could not find benefit")
		return nil, err
	}
	deleteBenefit.Status = benefit.DELETED 
	err = s.benefitRepo.Update(ctx, deleteBenefit)
	if err != nil {
		fmt.Println("Could not update changes:", err)
		return nil, err
	}
	benefitDTO := benefit.NewBenefitInfoDTO(*deleteBenefit)

	return benefitDTO, nil
}
