package services

import (
	"context"
	"fmt"

	"github.com/orbit-alliance/orbit-backend/internal/domain/nft"
	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
)

type Register42Service struct {
	userRepo     user.UserRepository
	eventBus     *shared.EventBus
	api42Gateway user.Api42Gateway
	ethGateway   user.BlockchainGateway
}

func NewRegister42Service(
	userRepo user.UserRepository,
	eventBus *shared.EventBus,
	api42Gateway user.Api42Gateway,
	ethGateway user.BlockchainGateway,
) *Register42Service {
	return &Register42Service{
		userRepo:     userRepo,
		eventBus:     eventBus,
		api42Gateway: api42Gateway,
		ethGateway:   ethGateway,
	}
}
func (s *Register42Service) Register42User(ctx context.Context, code42, walletAddress string) (*user.UserBasicInfoDTO, error) {

	err := s.checkIfWalletAlreadyRegistered(ctx, walletAddress)
	if err != nil {
		fmt.Println("Wallet already registered:", err)
		return nil, err
	}

	token42, err := s.api42Gateway.ExchangeCodeForToken(ctx, code42)
	if err != nil {
		fmt.Println("Error exchanging code for token:", err)
		return nil, err
	}

	basicInfo, err := s.api42Gateway.GetBasicUserInfo(ctx, token42)
	if err != nil {
		fmt.Println("Error getting basic user info:", err)
		return nil, err
	}

	usr, err := s.userRepo.FindByID42(ctx, basicInfo.ID42)
	if err != nil {
		fmt.Println("Error finding user by ID42:", err)
		return nil, err
	}

	if usr != nil {
		err := s.changeWalletAddress(ctx, usr, walletAddress)
		if err != nil {
			fmt.Println("Error changing wallet address:", err)
			return nil, err
		}
		return basicInfo, err
	}

	coinStatus, err := s.ethGateway.GetCoinsStatusByWallet(ctx, walletAddress)
	if err != nil {
		fmt.Println("Error getting coin status by wallet:", err)
		return nil, err
	}

	usr = user.NewUser(shared.NewMongoID(), basicInfo.ID42, walletAddress, basicInfo.Login, *coinStatus, []nft.NFT{})
	err = s.userRepo.Save(ctx, usr)
	if err != nil {
		fmt.Println("Error saving new user:", err)
		return nil, err
	}

	s.eventBus.Publish(user.NewUser42Registered(usr))

	return basicInfo, nil
}

func (s *Register42Service) checkIfWalletAlreadyRegistered(ctx context.Context, wallet string) error {
	usrWallet, err := s.userRepo.FindByWallet(ctx, wallet)
	if err != nil {
		return err
	}
	if usrWallet != nil {
		return user.ErrWalletAlreadyRegistered
	}
	return nil
}

func (s *Register42Service) changeWalletAddress(ctx context.Context, usr *user.User, newWallet string) error {
	evnt, err := usr.ChangeWalletAddress(newWallet, usr.CoinStatus)
	if err != nil {
		return err
	}

	if err := s.userRepo.Save(ctx, usr); err != nil {
		return err
	}

	s.eventBus.Publish(evnt)

	return nil
}
