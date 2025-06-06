package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/orbit-alliance/orbit-backend/internal/application/user/services"
	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
	"github.com/orbit-alliance/orbit-backend/internal/infra/db"
	repo "github.com/orbit-alliance/orbit-backend/internal/infra/db/repos"
	"github.com/orbit-alliance/orbit-backend/internal/infra/gateway_42"
	"github.com/orbit-alliance/orbit-backend/internal/infra/web3"

	seed "github.com/orbit-alliance/orbit-backend/internal/infra/db/seeds"

	usercontroller "github.com/orbit-alliance/orbit-backend/internal/interface/http/controllers/user"
	"github.com/orbit-alliance/orbit-backend/internal/interface/http/routes"
)

func main() {

	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env", err)
	}

	// Obtém a porta da aplicação a partir das variáveis de ambiente
	APP_PORT := os.Getenv("APP_PORT")
	fmt.Println("Porta da aplicação: ", APP_PORT)

	// Inicializa o banco de dados MongoDB
	db.InitMongoDB()
	defer db.CloseMongoDB()

	// EventBus e gateways, criados para gerenciar eventos e interações com a blockchain
	eventBus := shared.NewEventBus()

	// Cria o EthGateway mockado
	ethGateway := web3.NewMockEthGateway()

	// Cria o EthGateway para interações com a blockchain Ethereum
	// ethGateway, err := web3.NewEthGateway(64)
	// if err != nil {
	// 	log.Fatal("Erro ao criar o EthGateway", err)
	// }

	// Cria o Api42Gateway para interações com a API da 42 de autenticação
	auth42Gateway := gateway_42.NewGateway42()

	// Repositórios para acessar os dados de usuários, transferências e ações de boas práticas
	userRepo := repo.NewUserRepository(db.Database)
	transferRepo := repo.NewTransferRepository(db.Database)
	goodActionRepo := repo.NewGoodActionRepository(db.Database)

	// Serviços que encapsulam a lógica de negócios
	transferCoinsService := services.NewTransferCoinsService(userRepo, transferRepo, eventBus)
	//register42Service := services.NewRegister42Service(userRepo, eventBus, auth42Gateway, ethGateway)

	// Inicializa o EthGateway para ouvir eventos de transferências
	ethGateway.TransferListener(context.TODO(), func(event user.TransferDTO) {
		if err := transferCoinsService.TransferCoins(event); err != nil {
			log.Printf("Failed to transfer coins for event %+v: %v", event, err)
		}
	})

	// Semente de dados iniciais para ações de boas práticas
	seed.SeedGoodActions(context.TODO(), goodActionRepo)

	userBonusProjectService := services.NewBonusProjectService(
		userRepo,
		eventBus,
		repo.NewUserProjectRepository(db.Database),
		repo.NewUserGoodActionRepository(db.Database),
		goodActionRepo,
		auth42Gateway,
	)

	// Publicador de ações de boas práticas na blockchain
	onChainPublisher := services.NewOnChainPublisher(ethGateway)
	eventBus.Subscribe(user.DidGoodAction{}.EventType(), onChainPublisher.GoodActionPublisher)
	// Inscreve o serviço de bônus de projetos para receber eventos de registro de usuários
	eventBus.Subscribe(user.User42Registered{}.EventType(), userBonusProjectService.ApplyRetroactiveBonusProject)

	// Handlers para as rotas de autenticação da 42
	auth42Handler := usercontroller.NewAuth42Handler(auth42Gateway)
	router := mux.NewRouter()
	routes.RegisterUserRoutes(router, auth42Handler)

	// Start HTTP server, para escutar na porta definida
	fmt.Println("API Orbit rodando em http://localhost:" + APP_PORT)
	log.Fatal(http.ListenAndServe(":"+APP_PORT, router))
}
