package services

import (
	"context"
	"fmt"
	"time"

	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	Benefit "github.com/orbit-alliance/orbit-backend/internal/domain/benefit"
	Store "github.com/orbit-alliance/orbit-backend/internal/domain/store"
	User "github.com/orbit-alliance/orbit-backend/internal/domain/user"
	Coin "github.com/orbit-alliance/orbit-backend/internal/domain/coin"
)

type BuyBenefitService struct {
	userRepo				User.UserRepository
	benefitRepo				Benefit.BenefitRepository
	storeRepo				Store.StoreRepository
	transferRepo			User.TransferRepository
	benefitPurchaseRepo		User.UserBenefitPurchaseRepository
	eventBus				*shared.EventBus
}

func NewBuyBenefitService(
	userRepo User.UserRepository, 
	benefitRepo Benefit.BenefitRepository,
	storeRepo Store.StoreRepository,
	transferRepo User.TransferRepository,
	benefitPurchaseRepo User.UserBenefitPurchaseRepository,
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

func (s *BuyBenefitService) BuyBenefit(userId, benefitId string) (*User.UserBenefitPurchaseDTO, error) {
	
	ctx := context.Background()

	buyer, err := s.userRepo.FindByID(ctx, userId)
	if err != nil {
		fmt.Println("User not found:", err)
		return nil, err
	}
	if buyer == nil {
		return nil, User.ErrSenderNotFound
	}

	benefit, err := s.benefitRepo.FindByID(ctx, benefitId)
	if err != nil {
		fmt.Println("Benefit not found:", err)
		return nil, err
	}
	if benefit == nil {
		return nil, Benefit.ErrBenefitNotFound
	}
	if benefit.TotalAvailable == 0 {
		fmt.Println("This benefit is sold out:", err)
		return nil, err
	}
	if benefit.Status != Benefit.ACTIVE {
		fmt.Println("This benefit cannot be bought:", err)
		return nil, err
	}

	store, err := s.storeRepo.GetSingle(ctx)
	if err != nil {
		fmt.Println("Store not found:", err)
		return nil, err
	}
	if store == nil {
		fmt.Println("Store not found:", err)
		return nil, Store.ErrStoreNotFound 
	}
	if store.Status == Store.INACTIVE {
		fmt.Println("Store is inactive:", err)
		return nil, Store.ErrStoreNotFound 
	}

	// TODO buying logics here
	evt, err := buyer.SendCoins(&store.AdmUser, benefit.EarnedCost)
	if err != nil {
		return nil, err
	}
	s.eventBus.Publish(evt)
	evt2, err := store.AdmUser.ReceiveCoins(buyer, benefit.EarnedCost)
	if err != nil {
		return nil, err
	}
	s.eventBus.Publish(evt2)

	if (store.BenefitQuantity != 0) {
		store.BenefitQuantity--;
	}
	if (benefit.TotalAvailable != 0) {
		benefit.TotalAvailable--
	}
	store.TimeLastSale = time.Now()

	err = s.userRepo.Save(ctx, buyer);
	if err != nil {
		fmt.Println("Could not update user:", err)
		return nil, err
	}
	err = s.storeRepo.Save(ctx, store)
	if err != nil {
		fmt.Println("Could not update store:", err)
		return nil, err
	}
	err = s.userRepo.Save(ctx, &store.AdmUser)
	if err != nil {
		fmt.Println("Could not update store:", err)
		return nil, err
	}

	transfer := Coin.NewTransfer(buyer.ID.Hex(), buyer.Username, store.AdmUser.ID.Hex(), store.AdmUser.Username, benefit.EarnedCost)
	err = s.transferRepo.Save(ctx, transfer)
	if err != nil {
		fmt.Println("Could not register coin transference:", err)
		return nil, err
	}

	benefitPurchase := User.NewUserBenefitPurchase(buyer.ID42, 
		buyer.Username, benefit.ID, benefit.Name, benefit.EarnedCost, benefit.TransferredCost)
	err = s.benefitPurchaseRepo.Save(ctx, benefitPurchase)
	if err != nil {
		fmt.Println("Could not finish purchase:", err)
		return nil, err
	}

	purchaseDTO := User.NewUserBenefitPurchaseDTO(*benefitPurchase)

	return purchaseDTO, nil
}
