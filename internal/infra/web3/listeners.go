package web3

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
)

func (e *EthGateway) TransferListener(ctx context.Context, handler func(event user.TransferDTO)) error {
	opts := &bind.WatchOpts{Context: ctx}

	transferChan := make(chan *GalactoTransfer)

	transferSub, err := e.contract.WatchTransfer(opts, transferChan, []common.Address{}, []common.Address{})
	if err != nil {
		return err
	}

	go func() {
		defer func() {
			transferSub.Unsubscribe()
		}()

		for {
			select {
			case <-ctx.Done():
				return

			case transfer := <-transferChan:
				if transfer != nil {
					handler(galactoTransferToDTO(transfer))
				}
			}
		}
	}()

	return nil
}
