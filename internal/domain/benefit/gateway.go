package benefit

type BenefitInfoDTO struct {
	ID				string		`json:"_id"`
	Name			string		`json:"name"`
	Description		string		`json:"description"` 
	EarnedCost		int64		`json:"earned_cost"`
	TransferedCost	int64		`json:"transfered_cost"`
	TotalAvailable	int64		`json:"total_available"`
	MaxPeruser		int64		`json:"max_per_user"`
	ImageURL		string		`json:"image_url"`
	Category		string		`json:"category"`
	Tags			[]string	`json:"tags"`
	IsActive		bool		`json:"is_active"`
	CreatedBy		string		`json:"created_by"`
	CreatedAt       string		`json:"created_at"`
	LastUpdated     string		`json:"last_updated"`
}

func NewBenefitInfoDTO(id, name, description, imageUrl, category, createdBy, createdAt, lastUpdated string, 
	earnedCost, transferedCost, totalAvailable, maxPerUser int64, isActive bool, tags []string) *BenefitInfoDTO {
	return &BenefitInfoDTO{
		ID:				id,
		Name:			name,
		Description:	description, 
		EarnedCost:		earnedCost,
		TransferedCost:	transferedCost,
		TotalAvailable:	totalAvailable,
		MaxPeruser:		maxPerUser,
		ImageURL:		imageUrl,
		Category:		category,
		Tags:			tags,
		IsActive:		isActive,
		CreatedBy:		createdBy,
		CreatedAt:      createdAt,
		LastUpdated:	lastUpdated,
	}
}
