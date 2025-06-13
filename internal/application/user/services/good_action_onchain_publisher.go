package services

import (
	"context"
	"fmt"

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
func (p *OnChainPublisher) GoodActionPublisher(ctx context.Context, event shared.DomainEvent) {

	if e, ok := event.(*user.DidGoodAction); ok {
		err := p.ethGateway.PublishUserAction(ctx, e.Action)
		if err != nil {
			fmt.Printf("Error publishing good action on chain: %v\n", err)
			return
		}
	}
}
