package routes

import (
	"github.com/gorilla/mux"
	usercontroller "github.com/orbit-alliance/orbit-backend/internal/interface/http/controllers/user"
)

func RegisterUserRoutes(router *mux.Router, userController *usercontroller.UserHandler) {
	router.HandleFunc("/api/v0/user", userController.Register42User).Methods("POST")
	router.HandleFunc("/api/v0/user_transfer", userController.TransferCoins).Methods("POST")
	router.HandleFunc("/api/v0/user_buy_benefit", userController.UserBuyBenefit).Methods("POST")
}
