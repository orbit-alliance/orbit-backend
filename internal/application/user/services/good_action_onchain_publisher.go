package services

import (
	"context"

	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
)

type OnChainPublisher struct {
	ethGateway user.BlockchainGateway
}

func NewOnChainPublisher(ethGateway user.BlockchainGateway) *OnChainPublisher {
	return &OnChainPublisher{
		ethGateway: ethGateway,
	}
}
func (p *OnChainPublisher) GoodActionPubliser(ctx context.Context, event shared.DomainEvent) {

	if e, ok := event.(*user.DidGoodAction); ok {
		p.ethGateway.PublishUserAction(ctx, e.Action)
	}
}
