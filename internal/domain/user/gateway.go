package user

import "context"

type BlockchainGateway interface {
	PublishUserAction(ctx context.Context, payload UserGoodAction) error
}
