package web3

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/orbit-alliance/orbit-backend/internal/domain/coin"
)

func (e *EthGateway) GetCoinsStatusByWallet(ctx context.Context, wallet string) (*coin.CoinStatus, error) {
	address := common.HexToAddress(wallet)

	callOpts := &bind.CallOpts{
		Context: ctx,
	}

	breakdown, err := e.contract.GetBalanceBreakdown(callOpts, address)
	if err != nil {
		return nil, err
	}

	return coin.NewCoinStatus(
		breakdown.FromStaff.Uint64(),
		breakdown.FromTransfer.Uint64(),
	), nil
}
