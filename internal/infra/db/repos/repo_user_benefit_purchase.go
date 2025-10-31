package repo

import (
	"context"

	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserBenefitPurchaseRepository struct {
	collection *mongo.Collection
}

func NewUserBenefitPurchaseRepository(db *mongo.Database) *UserBenefitPurchaseRepository {
	collection := db.Collection("user_benefit_purchases")
	return &UserBenefitPurchaseRepository{collection: collection}
}

func (r *UserBenefitPurchaseRepository) Save(ctx context.Context, purchase *user.UserBenefitPurchase) error {

	filter := bson.M{"_id": purchase.ID}
	update := bson.M{
		"$set": purchase,
	}

	opts := options.Update().SetUpsert(true)

	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}
