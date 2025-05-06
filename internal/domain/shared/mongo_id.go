package shared

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GenerateMongoID generates a new random ObjectID to be used as the ID
func NewMongoID() primitive.ObjectID {
	return primitive.NewObjectID()
}

// ObjectIDFromString converts a string to an ObjectID
func ObjectIDFromString(s string) (primitive.ObjectID, error) {
	objectID, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		log.Printf("Erro ao converter string para ObjectID: %v", err)
		return primitive.ObjectID{}, err
	}
	return objectID, nil
}

// ObjectIDToString converts an ObjectID to a string, if the ObjectID is empty, returns an empty string
func ObjectIDToString(o primitive.ObjectID) string {
	return o.Hex()
}
