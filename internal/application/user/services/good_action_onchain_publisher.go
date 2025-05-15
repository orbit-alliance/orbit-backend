package services

import (
	"context"

	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
	"github.com/orbit-alliance/orbit-backend/internal/infra/web3"
)

type OnChainPublisher struct {
	ethGateway web3.EthGateway
}

func NewOnChainPublisher(ethGateway web3.EthGateway) *OnChainPublisher {
	return &OnChainPublisher{
		ethGateway: ethGateway,
	}
}
func (p *OnChainPublisher) Handler(ctx context.Context, event shared.DomainEvent) {

	if e, ok := event.(*user.DidGoodAction); ok {
		p.ethGateway.PublishUserAction(ctx, e.Action)
	}
}
