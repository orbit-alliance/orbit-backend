package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	Benefit "github.com/orbit-alliance/orbit-backend/internal/domain/benefit"
	Store "github.com/orbit-alliance/orbit-backend/internal/domain/store"
	User "github.com/orbit-alliance/orbit-backend/internal/domain/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterBenefitService struct {
	benefitRepo  Benefit.BenefitRepository
	storeRepo  Store.StoreRepository
	userRepo	User.UserRepository
	eventBus   *shared.EventBus
}

func NewRegisterBenefitService(
	benefitRepo	Benefit.BenefitRepository,
	storeRepo	Store.StoreRepository,
	userRepo	User.UserRepository,
	eventBus	*shared.EventBus,
) *RegisterBenefitService {
	return &RegisterBenefitService{
		benefitRepo:	benefitRepo,
		storeRepo:		storeRepo,
		userRepo:		userRepo,
		eventBus:		eventBus,
	}
}

func isValidStoreAdmUser(s *RegisterBenefitService, ctx context.Context, store Store.Store, admUserId string) bool {
	user, err := s.userRepo.FindByID(ctx, admUserId)
	if err != nil || user == nil {
		fmt.Println("Store adm user not found:", err)
		return false
	}
	if (store.AdmUser.ID != user.ID) {
		fmt.Println("User is not a Store administrator:")
		return false
	}
	return true
}

func (s *RegisterBenefitService) RegisterBenefit(
	ctx context.Context, payload Benefit.BenefitPayload) (*Benefit.BenefitInfoDTO, error) {
	
	store, err := s.storeRepo.FindByID(ctx, payload.StoreId)
	if err != nil {
		fmt.Println("Store to register benefit not found:", err)
		return nil, err
	}
	if (!isValidStoreAdmUser(s, ctx, *store, payload.AdmUserId)) {
		return nil, errors.New("Could not validate user:")
	}

	newBenefit := Benefit.NewBenefit(primitive.NewObjectID(), payload, store.AdmUser.Username)
	err = s.benefitRepo.Save(ctx, newBenefit)
	if (err != nil) {
		fmt.Println("Error trying to save benefit:", err)
		return nil, err
	}
	store.BenefitQuantity += newBenefit.TotalAvailable
	err = s.storeRepo.Save(ctx, store)
	if (err != nil) {
		fmt.Println("Error trying to update store benefit quantity:", err)
		return nil, err
	}

	benefitDto := Benefit.NewBenefitInfoDTO(*newBenefit)
	s.eventBus.Publish(Benefit.NewBenefitRegistered(newBenefit))
	return benefitDto, nil 
}

func (s *RegisterBenefitService) UpdateBenefit(
	ctx context.Context, payload Benefit.BenefitPayload) (*Benefit.BenefitInfoDTO, error) {

	store, err := s.storeRepo.FindByID(ctx, payload.StoreId)
	if err != nil {
		fmt.Println("Store to register benefit not found:", err)
		return nil, err
	}
	if (!isValidStoreAdmUser(s, ctx, *store, payload.AdmUserId)) {
		return nil, errors.New("Could not validate user")
	}

	benefit, err := s.benefitRepo.FindByID(ctx, payload.BenefitId)
	benefit.Name = payload.Name
	benefit.Description = payload.Description
	benefit.ImageURL = payload.ImageURL
	benefit.Category = payload.Category
	benefit.EarnedCost = payload.EarnedCost
	benefit.TransferredCost = payload.TransferredCost
	benefit.MaxPerUser = payload.MaxPerUser
	benefit.LastUpdated = time.Now()
	benefit.Status = payload.Status
	benefit.Tags = payload.Tags
	store.BenefitQuantity -= benefit.TotalAvailable
	benefit.TotalAvailable = payload.TotalAvailable
	store.BenefitQuantity += benefit.TotalAvailable

	// TODO check if use Save or Update
	err = s.benefitRepo.Save(ctx, benefit)
	if err != nil {
		fmt.Println("Could not update Benefit changes:", err)
		return nil, err
	}

	err = s.storeRepo.Save(ctx, store)
	if err != nil {
		fmt.Println("Could not Store changes:", err)
		return nil, err
	}

	benefitDTO := Benefit.NewBenefitInfoDTO(*benefit)

	return benefitDTO, nil
}

func (s *RegisterBenefitService) DeleteBenefit(ctx context.Context, payload Benefit.BenefitPayload) (*Benefit.BenefitInfoDTO, error) {
	
	store, err := s.storeRepo.FindByID(ctx, payload.StoreId)
	if err != nil {
		fmt.Println("Store to register benefit not found:", err)
		return nil, err
	}
	if (!isValidStoreAdmUser(s, ctx, *store, payload.AdmUserId)) {
		return nil, errors.New("Could not validate user")
	}

	benefit, err := s.benefitRepo.FindByID(ctx, payload.BenefitId)
	if err != nil {
		fmt.Println("Could not find benefit")
		return nil, err
	}
	benefit.Status = Benefit.DELETED
	store.BenefitQuantity -= benefit.TotalAvailable
	benefit.TotalAvailable = 0
	// TODO check if use Save or Update
	err = s.benefitRepo.Save(ctx, benefit)
	if err != nil {
		fmt.Println("Could not update benefit changes:", err)
		return nil, err
	}
	
	err = s.storeRepo.Save(ctx, store)
	if err != nil {
		fmt.Println("Could not update store changes:", err)
		return nil, err
	}

	benefitDTO := Benefit.NewBenefitInfoDTO(*benefit)

	return benefitDTO, nil
}
