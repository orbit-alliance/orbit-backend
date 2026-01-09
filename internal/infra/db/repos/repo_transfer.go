package repo

import (
	"context"

	"github.com/orbit-alliance/orbit-backend/internal/domain/coin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TransferRepository struct {
	collection *mongo.Collection
}

func NewTransferRepository(db *mongo.Database) *TransferRepository {
	collection := db.Collection("transfers")
	return &TransferRepository{collection: collection}
}

func (r *TransferRepository) Save(ctx context.Context, transfer *coin.Transfer) error {
	filter := bson.M{"_id": transfer.ID}
	update := bson.M{
		"$set": transfer,
	}

	opts := options.Update().SetUpsert(true)

	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}
