package shared

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GenerateMongoID gera um novo ObjectID aleatório para ser usado como ID
func NewMongoID() primitive.ObjectID {
	return primitive.NewObjectID()
}

// ObjectIDFromString converte uma string para um ObjectID
func ObjectIDFromString(s string) (primitive.ObjectID, error) {
	objectID, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		log.Printf("Erro ao converter string para ObjectID: %v", err)
		return primitive.ObjectID{}, err
	}
	return objectID, nil
}

// ObjectIDToString converte um ObjectID para uma string, se o ObjectID for vazio, retorna uma string vazia
func ObjectIDToString(o primitive.ObjectID) string {
	return o.Hex()
}
