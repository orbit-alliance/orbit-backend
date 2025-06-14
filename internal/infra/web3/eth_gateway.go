package web3

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
)

type EthGateway struct {
	client          *ethclient.Client
	contract        *Galacto
	contractAddress common.Address

	pk      *ecdsa.PrivateKey
	chainID *big.Int

	queue chan *user.UserGoodAction
}

func NewEthGateway() (*EthGateway, error) {
	rpcURL := os.Getenv("ETH_RPC_URL") // Ex: http://127.0.0.1:8545
	pkHex := strings.TrimPrefix(os.Getenv("ETH_PRIVATE_KEY"), "0x")
	contractAddr := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RPC: %w", err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	pk, err := crypto.HexToECDSA(pkHex)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %w", err)
	}

	galacto, err := NewGalacto(contractAddr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to load contract: %w", err)
	}

	g := &EthGateway{
		client:          client,
		contract:        galacto,
		contractAddress: contractAddr,
		pk:              pk,
		chainID:         chainID,
		queue:           make(chan *user.UserGoodAction, 1_000),
	}

	go g.worker() // inicia o processamento em background

	return g, nil
}

// Método público que garante a publicação (bloqueia se fila cheia)
func (g *EthGateway) PublishUserAction(_ context.Context, action *user.UserGoodAction) {
	g.queue <- action
}

// Worker infinito com retry
func (g *EthGateway) worker() {
	for action := range g.queue {
		for {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			err := g.sendTx(ctx, action)
			cancel()

			if err == nil {
				break
			}

			log.Printf("⚠️ erro ao enviar tx: %v. Repetindo em 10s...", action.ID)
			log.Printf("Detalhes: %v", err)
			time.Sleep(10 * time.Second)
		}
	}
}

// Envio da transação + espera do mining
func (g *EthGateway) sendTx(ctx context.Context, a *user.UserGoodAction) error {
	auth, err := bind.NewKeyedTransactorWithChainID(g.pk, g.chainID)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %w", err)
	}
	auth.Context = ctx
	auth.GasLimit = 300_000 // ou deixe zero para estimar automaticamente

	tx, err := g.contract.PublishAction(
		auth,
		shared.ObjectIDToString(a.ID),
		common.HexToAddress(a.UserWallet),
		shared.ObjectIDToString(a.UserID),
		a.Username,
		shared.ObjectIDToString(a.ActionID),
		a.ActionName,
		big.NewInt(int64(a.RewardAmount)),
		big.NewInt(a.PerformedAt.Unix()),
	)
	if err != nil {
		return fmt.Errorf("failed to send tx: %w", err)
	}

	receipt, err := bind.WaitMined(ctx, g.client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for mining: %w", err)
	}
	if receipt.Status != 1 {
		return fmt.Errorf("transaction reverted: %s", tx.Hash().Hex())
	}

	log.Printf("✅ ação publicada com sucesso! txHash: %s, value: %d", tx.Hash().Hex(), a.RewardAmount)
	return nil
}
