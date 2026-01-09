package benefit

import "github.com/orbit-alliance/orbit-backend/internal/domain/shared"

type BenefitRegistered struct {
	shared.BaseEvent
	Benefit *Benefit
}

func (b BenefitRegistered) EventType() string {
	return "benefit.BenefitRegistered"
}

func NewBenefitRegistered(benefit *Benefit) *BenefitRegistered {
	return &BenefitRegistered {
		BaseEvent:	shared.NewBaseEvent(shared.NewMongoID()),
		Benefit:	benefit,
	}
}
