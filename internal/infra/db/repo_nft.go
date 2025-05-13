package db

import (
	"context"

	"github.com/orbit-alliance/orbit-backend/internal/domain/nft"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NFTRepository struct {
	collection *mongo.Collection
}

func NewNFTRepository(db *mongo.Database) *NFTRepository {
	collection := db.Collection("NFts")
	return &NFTRepository{collection: collection}
}

func (r *NFTRepository) Save(ctx context.Context, nft *nft.NFT) error {

	filter := bson.M{"_id": nft.ID}
	update := bson.M{
		"$set": nft,
	}

	opts := options.Update().SetUpsert(true)

	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

func (r *NFTRepository) LoadAll(ctx context.Context) []nft.NFT {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil
	}
	defer cursor.Close(ctx)

	var NFTs []nft.NFT
	for cursor.Next(ctx) {
		var goodAction nft.NFT
		if err := cursor.Decode(&goodAction); err != nil {
			return nil
		}
		NFTs = append(NFTs, goodAction)
	}

	if err := cursor.Err(); err != nil {
		return nil
	}

	return NFTs
}
