package repo

import (
	"context"
	"sync"

	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	collection *mongo.Collection
	lockMap    sync.Map
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	collection := db.Collection("users")
	return &UserRepository{collection: collection}
}

func (r *UserRepository) getMutex(id primitive.ObjectID) *sync.Mutex {
	actual, _ := r.lockMap.LoadOrStore(id, &sync.Mutex{})
	return actual.(*sync.Mutex)
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

func (r *UserRepository) EarnTokens(ctx context.Context, id primitive.ObjectID, amount uint64) error {

	filter := bson.M{"_id": id}
	update := bson.M{
		"$inc": bson.M{"coin_status.earned_by_actions": amount},
	}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) IncrementRewardedDays(ctx context.Context, id primitive.ObjectID, increment int) error {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$inc": bson.M{"rewarded_days": increment},
	}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) SaveLogin(ctx context.Context, id primitive.ObjectID, lastLoginIn42 string, currentStreak int) error {

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"last_login_in_42": lastLoginIn42,
			"current_streak":   currentStreak,
		},
	}
	_, err := r.collection.UpdateOne(ctx, filter, update)
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
	mu := r.getMutex(idObj)
	mu.Lock()
	defer mu.Unlock()

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

func (r *UserRepository) LoadAll(ctx context.Context) ([]*user.User, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*user.User
	for cursor.Next(ctx) {
		var user user.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
