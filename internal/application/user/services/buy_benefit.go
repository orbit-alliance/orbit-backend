package services

import (
	"context"
	"fmt"
	"time"

	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	Benefit "github.com/orbit-alliance/orbit-backend/internal/domain/benefit"
	Store "github.com/orbit-alliance/orbit-backend/internal/domain/store"
	User "github.com/orbit-alliance/orbit-backend/internal/domain/user"
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

func getUser(s *BuyBenefitService, ctx context.Context, userId string) (*User.User, error) {
	user, err := s.userRepo.FindByID(ctx, userId)
	if err != nil {
		fmt.Println("User not found:", err)
		return nil, err
	}
	if user == nil {
		return nil, User.ErrSenderNotFound
	}
	return user, nil
}

func getBenefit(s *BuyBenefitService, ctx context.Context, benefitId string) (*Benefit.Benefit, error) {
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
	return benefit, nil
}

func getStore(s *BuyBenefitService, ctx context.Context) (*Store.Store, error) {
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
	return store, nil
}

func (s *BuyBenefitService) BuyBenefit(userId, benefitId string) (*User.UserBenefitPurchaseDTO, error) {
	
	ctx := context.Background()

	user, err := getUser(s, ctx, userId)
	if err != nil { return nil, err }
	if user == nil { return nil, nil }

	benefit, err := getBenefit(s, ctx, benefitId)
	if err != nil { return nil, err }
	if benefit == nil { return nil, nil}
	
	store, err := getStore(s, ctx)
	if err != nil { return nil, err }
	if store == nil { return nil, nil }

	evt, err  := user.PurchaseBenefit(&store.AdmUser, benefit)
	if err != nil {
		return nil, err
	}
	s.eventBus.Publish(evt);

	// Use user.receivedCoins 
	evt2, err := store.ReceivePurchaseCoins(user, benefit)
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

	err = s.userRepo.Save(ctx, user);
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
	err = s.benefitRepo.Save(ctx, benefit)
	if err != nil {
		fmt.Println("Could not update benefit:", err)
		return nil, err
	}
	
	benefitPurchase := User.NewUserBenefitPurchase(user, benefit)
	err = s.benefitPurchaseRepo.Save(ctx, benefitPurchase)
	if err != nil {
		fmt.Println("Could not finish purchase:", err)
		return nil, err
	}

	purchaseDTO := User.NewUserBenefitPurchaseDTO(*benefitPurchase)

	return purchaseDTO, nil
}
