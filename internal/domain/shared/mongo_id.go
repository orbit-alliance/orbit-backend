package shared

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewMongoID() primitive.ObjectID {
	return primitive.NewObjectID()
}

func ObjectIDFromString(s string) (primitive.ObjectID, error) {
	objectID, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		log.Printf("Erro ao converter string para ObjectID: %v", err)
		return primitive.ObjectID{}, err
	}
	return objectID, nil
}

func ObjectIDToString(o primitive.ObjectID) string {
	return o.Hex()
}
