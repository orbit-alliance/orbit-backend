package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	user_controller "github.com/orbit-alliance/orbit-backend/internal/interface/http/controllers/user"
	store_controller "github.com/orbit-alliance/orbit-backend/internal/interface/http/controllers/store"
	"github.com/orbit-alliance/orbit-backend/internal/interface/http/routes"
)

func buildHTTP(cfg Config, userHandler *user_controller.UserHandler, storeHandler *store_controller.StoreHandler) *http.Server {
	r := mux.NewRouter()
	routes.RegisterUserRoutes(r, userHandler)
	routes.RegisterStoreRoutes(r, storeHandler)

	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{cfg.ALLOWED_ORIGINS}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	srv := &http.Server{
		Addr:    ":" + cfg.APP_PORT,
		Handler: cors(r),
	}

	return srv
}
