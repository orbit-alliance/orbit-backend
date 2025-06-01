package services

import (
	"context"

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
func (s *Register42Service) Register42User(ctx context.Context, code42, walletAddress string) (string, string, error) {

	err := s.checkIfWalletAlreadyRegistered(ctx, walletAddress)
	if err != nil {
		return "", "", err
	}

	token42, err := s.api42Gateway.ExchangeCodeForToken(ctx, code42)
	if err != nil {
		return "", "", err
	}

	ID42, login, err := s.api42Gateway.GetBasicUserInfo(ctx, token42)
	if err != nil {
		return "", "", err
	}

	usr, err := s.userRepo.FindByID42(ctx, ID42)
	if err != nil {
		return "", "", err
	}

	if usr != nil {
		err := s.changeWalletAddress(ctx, usr, walletAddress)
		return ID42, login, err
	}

	coinStatus, err := s.ethGateway.GetCoinsStatusByWallet(ctx, walletAddress)
	if err != nil {
		return "", "", err
	}

	usr = user.NewUser(shared.NewMongoID(), ID42, walletAddress, login, *coinStatus, []nft.NFT{})
	err = s.userRepo.Save(ctx, usr)
	if err != nil {
		return "", "", err
	}

	s.eventBus.Publish(ctx, user.NewUser42Registered(usr))

	return ID42, login, nil
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
	s.eventBus.Publish(ctx, evnt)

	return nil
}
