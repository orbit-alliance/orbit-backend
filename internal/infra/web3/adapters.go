package web3

import "github.com/orbit-alliance/orbit-backend/internal/domain/user"

func galactoTransferToDTO(transfer *GalactoTransfer) user.TransferDTO {
	return user.TransferDTO{
		From:   transfer.From.Hex(),
		To:     transfer.To.Hex(),
		Amount: uint64(transfer.Value.Int64()),
	}
}
