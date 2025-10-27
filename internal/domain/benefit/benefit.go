package benefit

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BenefitStatus uint8

const (
	ACTIVE BenefitStatus = iota
	INACTIVE
	DELETED
)

type Benefit struct {
	ID				primitive.ObjectID	`json:"id" bson:"_id"`
	Name			string				`json:"name" bson:"name"`
	Description		string				`json:"description" bson:"description"` 
	EarnedCost		uint64				`json:"earned_cost" bson:"earned_cost"`
	TransferedCost	uint64				`json:"transfered_cost" bson:"transfered_cost"`
	TotalAvailable	uint64				`json:"total_available" bson:"total_available"`
	MaxPerUser		uint64				`json:"max_per_user" bson:"max_per_user"`
	ImageURL		string				`json:"image_url" bson:"image_url"`
    Category		string				`json:"category" bson:"category"`
    Tags			[]string			`json:"tags" bson:"tags"`
    Status			BenefitStatus		`json:"status" bson:"status"`
    CreatedBy		string				`json:"created_by" bson:"created_by"`
	CreatedAt		time.Time			`json:"created_at" bson:"created_at"`
	LastUpdated		time.Time			`json:"last_updated" bson:"last_updated"`
	AvailableStart	time.Time			`json:"available_start" bson:"available_start"`
	AvailableEnd	time.Time			`json:"available_end" bson:"available_end"`
	DeletedTime		time.Time			`json:"deleted_time" bson:"deleted_time"` 
}

func NewBenefit (id primitive.ObjectID, name, description, imageUrl, category, createdBy string, 
	totalAvailable, maxPeruser uint64, status BenefitStatus, tags []string) *Benefit {
	return &Benefit{
		ID:				id,
		Name:			name,
		Description:	description,
		ImageURL:		imageUrl,
		Category:		category,
		CreatedBy:		createdBy,
		TotalAvailable: totalAvailable,
		MaxPerUser:		maxPeruser,
		Status:			status,
		CreatedAt:		time.Now(),
		LastUpdated:	time.Now(),
		AvailableStart:	time.Now(),
		AvailableEnd:	time.Now(),
		DeletedTime:	time.Time{},
		Tags:			tags,
	}
}
