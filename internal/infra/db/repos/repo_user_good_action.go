package repo

import (
	"context"

	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserGoodActionRepository struct {
	collection *mongo.Collection
}

func NewUserGoodActionRepository(db *mongo.Database) *UserGoodActionRepository {
	collection := db.Collection("user_good_actions")
	return &UserGoodActionRepository{collection: collection}
}

func (r *UserGoodActionRepository) Save(ctx context.Context, userGoodAction *user.UserGoodAction) error {

	filter := bson.M{"_id": userGoodAction.ID}
	update := bson.M{
		"$set": userGoodAction,
	}

	opts := options.Update().SetUpsert(true)

	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}
