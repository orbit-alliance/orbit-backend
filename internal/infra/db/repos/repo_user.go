package repo

import (
	"context"

	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	collection := db.Collection("Users")
	return &UserRepository{collection: collection}
}

func (r *UserRepository) Save(ctx context.Context, user *user.User) error {

	filter := bson.M{"_id": user.ID}
	update := bson.M{
		"$set": user,
	}

	opts := options.Update().SetUpsert(true)

	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*user.User, error) {

	idObj, err := shared.ObjectIDFromString(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": idObj}
	var user user.User
	err = r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByID42(ctx context.Context, id42 string) (*user.User, error) {
	filter := bson.M{"id_42": id42}
	var user user.User
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByWallet(ctx context.Context, wallet string) (*user.User, error) {
	filter := bson.M{"wallet": wallet}
	var user user.User
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*user.User, error) {
	filter := bson.M{"username": username}
	var user user.User
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
