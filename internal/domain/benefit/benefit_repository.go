package benefit

import "context"

type BenefitRepository interface {
	Save(ctx context.Context, benefit *Benefit) error
	FindByID(ctx context.Context, id string) (*Benefit, error)
	LoadAll(ctx context.Context) ([]*Benefit, error)
}
