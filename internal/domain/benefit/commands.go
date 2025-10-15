package benefit

func (b *Benefit) RegisterBenefit(benefit *Benefit) (*BenefitRegistered, error) {
	return NewBenefitRegistered(benefit), nil
}
