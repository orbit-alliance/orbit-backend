package benefit

import (
	"time"

	"github.com/orbit-alliance/orbit-backend/internal/domain/nft"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Benefit struct {
	ID				primitive.ObjectID	`json:"id" bson:"_id"`
	Name			string				`json:"name" bson:"name"`
	Description		string				`json:"description" bson:"description"` 
	EarnedCost		uint64				`json:"earned_cost" bson:"earned_cost"`
	TransferedCost	uint64				`json:"transfered_cost" bson:"transfered_cost"`
	TotalAvailable	uint64				`json:"total_available" bson:"total_available"`
	MaxPeruser		uint64				`json:"max_per_user" bson:"max_per_user"`
	AllowedNFTTypes	[]nft.NFT			`json:"allowed_nft_types" bson:"allowed_nft_types"`
	BlockedNFTTypes	[]nft.NFT			`json:"blocked_nft_types" bson:"blocked_nft_types"`
	ImageURL		string				`json:"image_url" bson:"image_url"`
    Category		string				`json:"category" bson:"category"`
    Tags			[]string			`json:"tags" bson:"tags"`
    IsActive		bool				`json:"is_active" bson:"is_active"`
    CreatedBy		string				`json:"created_by" bson:"created_by"`
	CreatedAt		time.Time			`json:"created_at" bson:"created_at"`
	LastUpdated		time.Time			`json:"last_updated" bson:"last_updated"`
}

func NewBenefit (id primitive.ObjectID, name, description, imageUrl, category, createdBy string, totalAvailable, maxPeruser uint64) *Benefit {
	return &Benefit{
		ID:				id,
		Name:			name,
		Description:	description,
		ImageURL:		imageUrl,
		Category:		category,
		CreatedBy:		createdBy,
		TotalAvailable: totalAvailable,
		MaxPeruser:		maxPeruser,
	}
}
