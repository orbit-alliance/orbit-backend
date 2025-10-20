package benefit

type BenefitInfoDTO struct {
	ID				string		`json:"_id"`
	Name			string		`json:"name"`
	Description		string		`json:"description"` 
	EarnedCost		uint64		`json:"earned_cost"`
	TransferedCost	uint64		`json:"transfered_cost"`
	TotalAvailable	uint64		`json:"total_available"`
	MaxPerUser		uint64		`json:"max_per_user"`
	ImageURL		string		`json:"image_url"`
	Category		string		`json:"category"`
	Tags			[]string	`json:"tags"`
	IsActive		bool		`json:"is_active"`
	CreatedBy		string		`json:"created_by"`
	CreatedAt       string		`json:"created_at"`
	LastUpdated     string		`json:"last_updated"`
}

func NewBenefitInfoDTO(id, name, description, imageUrl, category, createdBy, createdAt, lastUpdated string, 
	earnedCost, transferedCost, totalAvailable, maxPerUser uint64, isActive bool, tags []string) *BenefitInfoDTO {
	return &BenefitInfoDTO{
		ID:				id,
		Name:			name,
		Description:	description, 
		EarnedCost:		earnedCost,
		TransferedCost:	transferedCost,
		TotalAvailable:	totalAvailable,
		MaxPerUser:		maxPerUser,
		ImageURL:		imageUrl,
		Category:		category,
		Tags:			tags,
		IsActive:		isActive,
		CreatedBy:		createdBy,
		CreatedAt:      createdAt,
		LastUpdated:	lastUpdated,
	}
}

func NewBenefitDTO(b *Benefit) *BenefitInfoDTO {
	return &BenefitInfoDTO{
		ID:				b.ID.Hex(),
		Name:			b.Name,
		Description:	b.Description, 
		EarnedCost:		b.EarnedCost,
		TransferedCost:	b.TransferedCost,
		TotalAvailable:	b.TotalAvailable,
		MaxPerUser:		b.MaxPerUser,
		ImageURL:		b.ImageURL,
		Category:		b.Category,
		Tags:			b.Tags,
		IsActive:		b.IsActive,
		CreatedBy:		b.CreatedBy,
		CreatedAt:      b.CreatedAt.String(),
		LastUpdated:	b.LastUpdated.String(),
	}
}
