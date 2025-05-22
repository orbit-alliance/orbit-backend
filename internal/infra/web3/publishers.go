package web3

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
)

func (g *EthGateway) PublishUserAction(ctx context.Context, payload *user.UserGoodAction) error {
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
