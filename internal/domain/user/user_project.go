package user

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	DefaultMaxScore = uint8(100)
)

type UserProject struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	MaxScore uint8              `json:"max_score" bson:"max_score"`
	UserID   primitive.ObjectID `json:"user_id" bson:"user_id"`
	UserName string             `json:"user_name" bson:"user_name"`
}

func NewUserProject(id primitive.ObjectID, name string, maxScore uint8, userID primitive.ObjectID, userName string) *UserProject {
	return &UserProject{
		ID:       id,
		Name:     name,
		MaxScore: maxScore,
		UserID:   userID,
		UserName: userName,
	}
}
