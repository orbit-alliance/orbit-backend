package routes

import (
	"github.com/gorilla/mux"

	storeController "github.com/orbit-alliance/orbit-backend/internal/interface/http/controllers/store"
)

func RegisterStoreRoutes(router *mux.Router, storeController *storeController.StoreHandler) {
	router.HandleFunc("/api/v0/store", storeController.RegisterStore).Methods("POST")
}
