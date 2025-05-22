package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/orbit-alliance/orbit-backend/internal/application/user/services"
	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
	"github.com/orbit-alliance/orbit-backend/internal/infra/db"
	"github.com/orbit-alliance/orbit-backend/internal/infra/web3"
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

	eventBus := shared.NewEventBus()
	ethGateway, err := web3.NewEthGateway(64)
	if err != nil {
		log.Fatal("Erro ao criar o EthGateway", err)
	}
	onChainPublisher := services.NewOnChainPublisher(ethGateway)
	eventBus.Subscribe(user.DidGoodAction{}.EventType(), onChainPublisher.Handler)

	goodActionRepo := db.NewGoodActionRepository(db.Database)
	userRepo := db.NewUserRepository(db.Database)
	transferRepo := db.NewTransferRepository(db.Database)
	transferCoinsService := services.NewTransferCoinsService(userRepo, transferRepo, eventBus)

	ethGateway.TransferListener(context.TODO(), func(event user.TransferDTO) {
		if err := transferCoinsService.TransferCoins(event); err != nil {
			log.Printf("Failed to transfer coins for event %+v: %v", event, err)
		}
	})

	db.SeedGoodActions(context.TODO(), goodActionRepo)

}
