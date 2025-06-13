package services

import (
	"context"

	"github.com/orbit-alliance/orbit-backend/internal/domain/coin"
	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
)

type TransferCoinsService struct {
	userRepo     user.UserRepository
	transferRepo user.TransferRepository
	eventBus     *shared.EventBus
}

func NewTransferCoinsService(userRepo user.UserRepository, transferRepo user.TransferRepository, eventBus *shared.EventBus) *TransferCoinsService {
	return &TransferCoinsService{
		userRepo:     userRepo,
		transferRepo: transferRepo,
		eventBus:     eventBus,
	}
}

// TransferCoins transfers coins from one user to another
func (s *TransferCoinsService) TransferCoins(transfer user.TransferDTO) error {

	ctx := context.Background()

	sender, err := s.userRepo.FindByWallet(ctx, transfer.From)
	if err != nil {
		return err
	}
	if sender == nil {
		return user.ErrSenderNotFound
	}

	receiver, err := s.userRepo.FindByWallet(ctx, transfer.To)
	if err != nil {
		return err
	}

	evt, err := sender.SendCoins(receiver, transfer.Amount)
	if err != nil {
		return err
	}
	s.eventBus.Publish(evt)

	evt2, err := receiver.ReceiveCoins(sender, transfer.Amount)
	if err != nil {
		return err
	}
	s.eventBus.Publish(evt2)

	if err := s.userRepo.Save(ctx, receiver); err != nil {
		return err
	}
	if err := s.userRepo.Save(ctx, sender); err != nil {
		return err
	}

	coinTransfer := coin.NewTransfer(shared.ObjectIDToString(sender.ID), sender.Username, shared.ObjectIDToString(receiver.ID), receiver.Username, transfer.Amount)
	if err := s.transferRepo.Save(ctx, coinTransfer); err != nil {
		return err
	}

	return nil

}
