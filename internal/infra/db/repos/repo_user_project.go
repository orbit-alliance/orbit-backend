package repo

import (
	"context"

	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserProjectRepository struct {
	collection *mongo.Collection
}

func NewUserProjectRepository(db *mongo.Database) *UserProjectRepository {
	collection := db.Collection("user_projects")
	return &UserProjectRepository{collection: collection}
}

func (r *UserProjectRepository) Save(ctx context.Context, userProject *user.UserProject) error {

	filter := bson.M{"_id": userProject.ID}
	update := bson.M{
		"$set": userProject,
	}

	opts := options.Update().SetUpsert(true)
	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserProjectRepository) LoadByUserID(ctx context.Context, userID string) ([]*user.UserProject, error) {

	id, err := shared.ObjectIDFromString(userID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"user_id": id}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var userProjects []*user.UserProject
	for cursor.Next(ctx) {
		var userProject user.UserProject
		if err := cursor.Decode(&userProject); err != nil {
			return nil, err
		}
		userProjects = append(userProjects, &userProject)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return userProjects, nil
}
