package web3

import (
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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
