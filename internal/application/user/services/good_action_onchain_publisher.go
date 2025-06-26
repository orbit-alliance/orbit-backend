package services

import (
	"context"

	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
)

type OnChainGoodActionPublisher struct {
	ethGateway user.BlockchainGateway
}

func NewOnChainGoodActionPublisher(ethGateway user.BlockchainGateway) *OnChainGoodActionPublisher {
	return &OnChainGoodActionPublisher{
		ethGateway: ethGateway,
	}
}
func (p *OnChainGoodActionPublisher) GoodActionPublisher(ctx context.Context, event shared.DomainEvent) {

	if e, ok := event.(*user.DidGoodAction); ok {
		p.ethGateway.PublishUserAction(ctx, e.Action)
	}
}
