package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Database *mongo.Database

func InitMongoDB() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env", err)
	}

	uri := os.Getenv("MONGO_URL")
	fmt.Println("URI do MongoDB:", uri)

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Erro ao conectar ao MongoDB:", err)
	}
	Client = client

	Database = client.Database("orbit_alliance")

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Erro ao testar a conexão:", err)
	}
	fmt.Println("Conectado ao MongoDB!")
}

func CloseMongoDB() {
	err := Client.Disconnect(context.Background())
	if err != nil {
		log.Fatal("Erro ao desconectar do MongoDB:", err)
	}
	fmt.Println("Desconectado do MongoDB.")
}
