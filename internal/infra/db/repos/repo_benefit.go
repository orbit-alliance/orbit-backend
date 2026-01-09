package repo

import (
	"context"
	"sync"

	"github.com/orbit-alliance/orbit-backend/internal/domain/benefit"
	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BenefitRepository struct {
	collection *mongo.Collection
	lockMap    sync.Map
}

func NewBenefitRepository(db *mongo.Database) *BenefitRepository {
	collection := db.Collection("benefits")
	return &BenefitRepository{collection: collection}
}

func (r *BenefitRepository) getMutex(id primitive.ObjectID) *sync.Mutex {
	actual, _ := r.lockMap.LoadOrStore(id, &sync.Mutex{})
	return actual.(*sync.Mutex)
}

func (r *BenefitRepository) Save(ctx context.Context, benefit *benefit.Benefit) error {

	filter := bson.M{"_id": benefit.ID}
	update := bson.M{
		"$set": benefit,
	}

	opts := options.Update().SetUpsert(true)

	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

func (r *BenefitRepository) Update(ctx context.Context, benefit *benefit.Benefit) error {

	filter := bson.M{"_id": benefit.ID}
	update := bson.M{
		"$set": benefit,
	}

	// opts := options.Update().SetUpsert(true)

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *BenefitRepository) FindByID(ctx context.Context, id string) (*benefit.Benefit, error) {
	idObj, err := shared.ObjectIDFromString(id)
	if err != nil {
		return nil, err
	}
	mu := r.getMutex(idObj)
	mu.Lock()
	defer mu.Unlock()

	filter := bson.M{"_id": idObj}
	var benefit benefit.Benefit
	err = r.collection.FindOne(ctx, filter).Decode(&benefit)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &benefit, nil
}

func (r *BenefitRepository) LoadAll(ctx context.Context) ([]*benefit.Benefit, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var benefits []*benefit.Benefit
	for cursor.Next(ctx) {
		var benefit benefit.Benefit
		if err := cursor.Decode(&benefit); err != nil {
			return nil, err
		}
		benefits = append(benefits, &benefit)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return benefits, nil
}
