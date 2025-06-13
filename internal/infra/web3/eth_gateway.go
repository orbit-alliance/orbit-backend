package web3

import (
	"context"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthGateway struct {
	client          *ethclient.Client
	contract        *Galacto
	auth            *bind.TransactOpts
	contractAddress common.Address
}

func NewEthGateway() (*EthGateway, error) {
	rpcURL := os.Getenv("ETH_RPC_URL") // http://127.0.0.1:8545
	pkHex := strings.TrimPrefix(os.Getenv("ETH_PRIVATE_KEY"), "0x")
	contractAddr := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	// --> converte o HEX para *ecdsa.PrivateKey
	pk, err := crypto.HexToECDSA(pkHex)
	if err != nil {
		return nil, err
	}

	// --> cria o signer já com o chain‑id correto
	auth, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
	if err != nil {
		return nil, err
	}

	galacto, err := NewGalacto(contractAddr, client)
	if err != nil {
		return nil, err
	}

	return &EthGateway{
		client:          client,
		contract:        galacto,
		auth:            auth,
		contractAddress: contractAddr,
	}, nil
}
