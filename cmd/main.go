package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/orbit-alliance/orbit-backend/internal/infra/db"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env", err)
	}

	APP_PORT := os.Getenv("APP_PORT")
	fmt.Println("Porta da aplicação: ", APP_PORT)

	db.InitMongoDB()
	defer db.CloseMongoDB()

}
