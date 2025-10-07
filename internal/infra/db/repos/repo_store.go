package repo

import (
	"context"
	"sync"

	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Maybe shared
type StoreRepository struct {
	collection *mongo.Collection
	lockMap    sync.Map
}

func NewStoreRepository(db *mongo.Database) *StoreRepository {
	collection := db.Collection("store")
	return &StoreRepository{collection: collection}
}

func (r *StoreRepository) getMutex(id primitive.ObjectID) *sync.Mutex {
	actual, _ := r.lockMap.LoadOrStore(id, &sync.Mutex{})
	return actual.(*sync.Mutex)
}

func (r *StoreRepository) Save(ctx context.Context, store *store.Store) error {

	filter := bson.M{"_id": store.ID}
	update := bson.M{
		"$set": store,
	}

	opts := options.Update().SetUpsert(true)

	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

func (r *StoreRepository) FindByID(ctx context.Context, id string) (*store.Store, error) {
	idObj, err := shared.ObjectIDFromString(id)
	if err != nil {
		return nil, err
	}
	mu := r.getMutex(idObj)
	mu.Lock()
	defer mu.Unlock()

	filter := bson.M{"_id": idObj}
	var store store.Store
	err = r.collection.FindOne(ctx, filter).Decode(&store)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &store, nil
}

func (r *StoreRepository) LoadAll(ctx context.Context) ([]*store.Store, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var stores []*store.Store
	for cursor.Next(ctx) {
		var store store.Store
		if err := cursor.Decode(&store); err != nil {
			return nil, err
		}
		stores = append(stores, &store)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return stores, nil
}
