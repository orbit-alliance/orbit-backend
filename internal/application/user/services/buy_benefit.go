package services

import (
	"context"
	"fmt"
	"time"

	"github.com/orbit-alliance/orbit-backend/internal/domain/benefit"
	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/store"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
)

type BuyBenefitService struct {
	userRepo				user.UserRepository
	benefitRepo				benefit.BenefitRepository
	storeRepo				store.StoreRepository
	transferRepo			user.TransferRepository
	benefitPurchaseRepo		user.UserBenefitPurchaseRepository
	eventBus				*shared.EventBus
}

func NewBuyBenefitService(
	userRepo user.UserRepository, 
	benefitRepo benefit.BenefitRepository,
	storeRepo store.StoreRepository,
	transferRepo user.TransferRepository,
	benefitPurchaseRepo user.UserBenefitPurchaseRepository,
	eventBus *shared.EventBus) *BuyBenefitService {
	return &BuyBenefitService{
		userRepo:	userRepo,
		benefitRepo: benefitRepo,
		storeRepo: storeRepo,
		transferRepo: transferRepo,
		benefitPurchaseRepo: benefitPurchaseRepo,
		eventBus: eventBus,
	}
}

func (s *BuyBenefitService) BuyBenefit(userId, benefitId string) (*user.UserBenefitPurchaseDTO, error) {
	
	ctx := context.Background()

	buyer, err := s.userRepo.FindByID(ctx, userId)
	if err != nil {
		fmt.Println("User not found:", err)
		return nil, err
	}
	if buyer == nil {
		return nil, user.ErrSenderNotFound
	}

	benefitSold, err := s.benefitRepo.FindByID(ctx, benefitId)
	if err != nil {
		fmt.Println("Benefit not found:", err)
		return nil, err
	}
	if benefitSold == nil {
		return nil, benefit.ErrBenefitNotFound
	}

	storeSeller, err := s.storeRepo.GetSingle(ctx)
	if err != nil {
		fmt.Println("Store not found:", err)
		return nil, err
	}
	if storeSeller == nil {
		fmt.Println("Store not found:", err)
		return nil, store.ErrStoreNotFound 
	}

	evt, err := buyer.SendCoins(&storeSeller.AdmUser, benefitSold.EarnedCost)
	if err != nil {
		return nil, err
	}
	s.eventBus.Publish(evt)
	evt2, err := storeSeller.AdmUser.ReceiveCoins(buyer, benefitSold.EarnedCost)
	if err != nil {
		return nil, err
	}
	s.eventBus.Publish(evt2)
	

	// storeSeller.BenefitQuantity--; // Need to calc
	storeSeller.TimeLastSale = time.Now()

	err = s.userRepo.Save(ctx, buyer);
	if err != nil {
		fmt.Println("Could not update user:", err)
		return nil, err
	}
	err = s.storeRepo.Save(ctx, storeSeller)
	if err != nil {
		fmt.Println("Could not update store:", err)
		return nil, err
	}

	benefitPurchase := user.NewUserBenefitPurchase(buyer.ID42, 
		buyer.Username, benefitSold.ID, benefitSold.Name, benefitSold.EarnedCost, benefitSold.TransferredCost)

	err = s.benefitPurchaseRepo.Save(ctx, benefitPurchase)
	if err != nil {
		fmt.Println("Could not finish purchase:", err)
		return nil, err
	}

	purchaseDTO := user.NewUserBenefitPurchaseDTO(*benefitPurchase)

	return purchaseDTO, nil
}
