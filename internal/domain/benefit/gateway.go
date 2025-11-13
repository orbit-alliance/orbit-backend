package benefit

import "time"

type BenefitInfoDTO struct {
	ID				string				`json:"_id"`
	Name			string				`json:"name"`
	Description		string				`json:"description"` 
	EarnedCost		uint64				`json:"earned_cost"`
	TransferredCost	uint64				`json:"transferred_cost"`
	TotalAvailable	uint64				`json:"total_available"`
	MaxPerUser		uint64				`json:"max_per_user"`
	ImageURL		string				`json:"image_url"`
	Category		string				`json:"category"`
	Tags			[]string			`json:"tags"`
	Status			BenefitStatus		`json:"status"`
	CreatedBy		string				`json:"created_by"`
	CreatedAt       time.Time			`json:"created_at"`
	LastUpdated     time.Time			`json:"last_updated"`
	AvailableStart	time.Time			`json:"available_start"`
	AvailableEnd	time.Time			`json:"available_end"`
}

func NewBenefitInfoDTO(benefit Benefit) *BenefitInfoDTO {
	return &BenefitInfoDTO{
		ID:					benefit.ID.Hex(),
		Name:				benefit.Name,
		Description:		benefit.Description, 
		EarnedCost:			benefit.EarnedCost,
		TransferredCost:	benefit.TransferredCost,
		TotalAvailable:		benefit.TotalAvailable,
		MaxPerUser:			benefit.MaxPerUser,
		ImageURL:			benefit.ImageURL,
		Category:			benefit.Category,
		Tags:				benefit.Tags,
		Status:				benefit.Status,
		CreatedBy:			benefit.CreatedBy,
		CreatedAt:			benefit.CreatedAt,
		LastUpdated:		benefit.LastUpdated,
		AvailableStart:		benefit.AvailableStart,
		AvailableEnd:		benefit.AvailableEnd,
	}
}

type BenefitPayload struct {
	StoreId			string					`json:"store_id"`
	AdmUserId		string					`json:"adm_user_id"`
	Name			string					`json:"name"`
	Description		string					`json:"description"`
	ImageURL		string					`json:"image_url"`
	Category		string					`json:"category"`
	EarnedCost		uint64					`json:"earned_cost"`
	TransferredCost	uint64					`json:"transferred_cost"`
	TotalAvailable	uint64					`json:"total_available"`
	MaxPerUser		uint64					`json:"max_per_user"`
	BenefitId		string					`json:"benefit_id"`
	Tags			[]string				`json:"tags"`	
	Status			BenefitStatus			`json:"status"`
}

