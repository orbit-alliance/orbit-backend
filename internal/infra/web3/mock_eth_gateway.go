package web3

import (
	"context"
	"log"

	"github.com/orbit-alliance/orbit-backend/internal/domain/coin"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
)

type MockEthGateway struct{}

func NewMockEthGateway() *MockEthGateway {
	return &MockEthGateway{}
}

func (m *MockEthGateway) PublishUserAction(ctx context.Context, payload *user.UserGoodAction) error {
	log.Println("🪄 [Mock] PublishUserAction chamado → payload:", payload)
	return nil
}

func (m *MockEthGateway) TransferListener(ctx context.Context, handler func(event user.TransferDTO)) error {
	log.Println("🪄 [Mock] TransferListener iniciado → ouvindo nada, apenas mock")
	return nil
}

func (m *MockEthGateway) GetCoinsStatusByWallet(ctx context.Context, wallet string) (*coin.CoinStatus, error) {
	log.Println("🪄 [Mock] GetCoinsStatusByWallet chamado para wallet:", wallet)
	// Retorna um estado de saldo fake
	return &coin.CoinStatus{
		EarnedByActions:  100,
		EarnedByTransfer: 50,
	}, nil
}
