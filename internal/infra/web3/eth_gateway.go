package web3

import (
	"context"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
)

type EthGateway struct {
	client          *ethclient.Client
	contract        *Galacto
	auth            *bind.TransactOpts
	contractAddress common.Address
}

func NewEthGateway(chainID int64) (*EthGateway, error) {
	rpcURL := os.Getenv("ETH_RPC_URL")
	privateKey := os.Getenv("ETH_PRIVATE_KEY")
	contractAddr := os.Getenv("CONTRACT_ADDRESS")

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewTransactorWithChainID(strings.NewReader(privateKey), "", big.NewInt(chainID))
	if err != nil {
		return nil, err
	}

	contract, err := NewGalacto(common.HexToAddress(contractAddr), client)
	if err != nil {
		return nil, err
	}

	return &EthGateway{
		client:          client,
		contract:        contract,
		auth:            auth,
		contractAddress: common.HexToAddress(contractAddr),
	}, nil
}

func (g *EthGateway) PublishUserAction(ctx context.Context, payload user.UserGoodAction) error {
	tx, err := g.contract.PublishAction(
		g.auth,
		shared.ObjectIDToString(payload.ID),
		common.HexToAddress(payload.UserWallet),
		shared.ObjectIDToString(payload.UserID),
		payload.Username,
		shared.ObjectIDToString(payload.ActionID),
		payload.ActionName,
		big.NewInt(int64(payload.RewardAmount)),
		big.NewInt(payload.PerformedAt.Unix()),
	)
	if err != nil {
		return err
	}

	log.Printf("Transaction submitted: %s", tx.Hash().Hex())
	return nil
}
