package db

import (
	"context"

	"github.com/orbit-alliance/orbit-backend/internal/domain/coin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GoodActionRepository struct {
	collection *mongo.Collection
}

func NewGoodActionRepository(db *mongo.Database) *GoodActionRepository {
	collection := db.Collection("good_actions")
	return &GoodActionRepository{collection: collection}
}

func (r *GoodActionRepository) Save(ctx context.Context, goodAction *coin.GoodAction) error {

	filter := bson.M{"_id": goodAction.ID}
	update := bson.M{
		"$set": goodAction,
	}

	opts := options.Update().SetUpsert(true)

	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

func (r *GoodActionRepository) LoadAll(ctx context.Context) []coin.GoodAction {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil
	}
	defer cursor.Close(ctx)

	var goodActions []coin.GoodAction
	for cursor.Next(ctx) {
		var goodAction coin.GoodAction
		if err := cursor.Decode(&goodAction); err != nil {
			return nil
		}
		goodActions = append(goodActions, goodAction)
	}

	if err := cursor.Err(); err != nil {
		return nil
	}

	return goodActions
}
