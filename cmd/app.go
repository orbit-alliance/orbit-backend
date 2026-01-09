package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	storeServices "github.com/orbit-alliance/orbit-backend/internal/application/store/services"
	benefitServices "github.com/orbit-alliance/orbit-backend/internal/application/benefit/services"
	"github.com/orbit-alliance/orbit-backend/internal/application/user/services"
	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
	"github.com/orbit-alliance/orbit-backend/internal/infra/db"
	repo "github.com/orbit-alliance/orbit-backend/internal/infra/db/repos"
	seed "github.com/orbit-alliance/orbit-backend/internal/infra/db/seeds"
	"github.com/orbit-alliance/orbit-backend/internal/infra/gateway_42"
	"github.com/orbit-alliance/orbit-backend/internal/infra/web3"
	store_controller "github.com/orbit-alliance/orbit-backend/internal/interface/http/controllers/store"
	user_controller "github.com/orbit-alliance/orbit-backend/internal/interface/http/controllers/user"
	benefit_controller "github.com/orbit-alliance/orbit-backend/internal/interface/http/controllers/benefit"
	"github.com/robfig/cron/v3"
)

type App struct {
	router *http.Server // ou *mux.Router, se preferir
	jobs   []shared.Job
	c      *cron.Cron
}

func NewApp(cfg Config) (*App, func()) {
	// ---------- Infra ----------
	db.InitMongoDB()
	eventBus := shared.NewEventBus()
	ethGateway, err := web3.NewEthGateway()
	if err != nil {
		panic("Failed to create EthGateway: " + err.Error())
	}
	// …

	// ---------- Cron ------------
	c := cron.New() // Inicia o cron com suporte a segundos

	// ---------- Gateways externos ----------
	gateway42 := gateway_42.NewGateway42()
	// …

	// ---------- Repositórios ----------
	userRepo := repo.NewUserRepository(db.Database)
	transferRepo := repo.NewTransferRepository(db.Database)
	goodActionRepo := repo.NewGoodActionRepository(db.Database)
	userProjectRepo := repo.NewUserProjectRepository(db.Database)
	userGoodActionRepo := repo.NewUserGoodActionRepository(db.Database)
	storeRepo := repo.NewStoreRepository(db.Database)
	benefitRepo := repo.NewBenefitRepository(db.Database)
	userBenefitPurchaseRepo := repo.NewUserBenefitPurchaseRepository(db.Database)
	// …

	// ---------- Seeds ----------
	seed.SeedGoodActions(context.TODO(), goodActionRepo)
	// …

	// ---------- Serviços ----------
	register42Service := services.NewRegister42Service(userRepo, eventBus, gateway42, ethGateway)
	transferCoinsService := services.NewTransferCoinsService(userRepo, transferRepo, eventBus)
	buyBenefitService := services.NewBuyBenefitService(userRepo, benefitRepo, storeRepo, transferRepo, userBenefitPurchaseRepo, eventBus)
	userBonusProjectService := services.NewBonusProjectService(
		userRepo,
		eventBus,
		userProjectRepo,
		userGoodActionRepo,
		goodActionRepo,
		gateway42,
	)
	userFrequencyRewardService := services.NewRetroactiveFrequencyRewardService(
		userRepo,
		goodActionRepo,
		userGoodActionRepo,
		gateway42,
		eventBus,
	)
	goodActionPublisher := services.NewOnChainGoodActionPublisher(ethGateway)
	registerStoreService := storeServices.NewRegisterStoreService(storeRepo, userRepo, eventBus, ethGateway)
	registerBenefitService := benefitServices.NewRegisterBenefitService(benefitRepo, storeRepo, userRepo, eventBus)

	// …

	// ---------- Controladores ----------
	userController := user_controller.NewUserHandler(register42Service, transferCoinsService, buyBenefitService)
	storeController := store_controller.NewStoreHandler(registerStoreService)
	benefitController := benefit_controller.NewBenefitHandler(registerBenefitService)
	// …

	// ---------- Register Domain Events ----------
	eventBus.Subscribe(user.DidGoodAction{}.EventType(), goodActionPublisher.GoodActionPublisher)
	eventBus.Subscribe(user.User42Registered{}.EventType(), userBonusProjectService.ApplyRetroactiveBonusProject)
	eventBus.Subscribe(user.User42Registered{}.EventType(), userFrequencyRewardService.ApplyRetroactiveFrequencyReward)
	// …

	// ---------- Escutadores Onchain ----------
	ethGateway.TransferListener(context.TODO(), func(transfer user.TransferDTO) {
		if err := transferCoinsService.TransferCoins(transfer); err != nil {
			log.Printf("Failed to transfer coins for event %+v: %v", transfer, err)
		}
	})
	// …

	// ---------- Jobs ----------
	jobs := []shared.Job{
		userBonusProjectService.DailyBonusProjectJob,
		userFrequencyRewardService.DailyFrequencyRewardJob,
	}

	// ---------- HTTP ----------
	router := buildHTTP(cfg, userController, storeController, benefitController /* , outros handlers ... */)

	app := &App{router: router, jobs: jobs, c: c}

	// cleanup
	cleanup := func() {
		db.CloseMongoDB()
		c.Stop() // Para o cron
		if err := app.router.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down HTTP server: %v", err)
		}
		log.Println("Application shutdown complete.")
	}

	return app, cleanup
}

func (a *App) Run(ctx context.Context) error {
	/* -------- inicia HTTP em goroutine -------- */
	go func() {
		if err := a.router.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("HTTP server error: %v", err)
		}
	}()
	log.Printf("HTTP server listening on %s", a.router.Addr)

	/* -------- inicia jobs em goroutines -------- */
	a.c.Start()
	dailyJobs(a.c, a.jobs)

	/* -------- espera sinal de término -------- */
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-quit:
		log.Println("Shutting down…")
	case <-ctx.Done():
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return a.router.Shutdown(shutdownCtx)
}
