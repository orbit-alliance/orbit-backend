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
	repo "github.com/orbit-alliance/orbit-backend/internal/infra/db/repos"
	seed "github.com/orbit-alliance/orbit-backend/internal/infra/db/seeds"
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

	goodActionRepo := repo.NewGoodActionRepository(db.Database)
	userRepo := repo.NewUserRepository(db.Database)
	transferRepo := repo.NewTransferRepository(db.Database)
	transferCoinsService := services.NewTransferCoinsService(userRepo, transferRepo, eventBus)

	ethGateway.TransferListener(context.TODO(), func(event user.TransferDTO) {
		if err := transferCoinsService.TransferCoins(event); err != nil {
			log.Printf("Failed to transfer coins for event %+v: %v", event, err)
		}
	})

	seed.SeedGoodActions(context.TODO(), goodActionRepo)

}
