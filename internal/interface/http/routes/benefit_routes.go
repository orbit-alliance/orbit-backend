package routes

import (
	"github.com/gorilla/mux"

	benefitController "github.com/orbit-alliance/orbit-backend/internal/interface/http/controllers/benefit"
)

func RegisterBenefitRoutes(router *mux.Router, benefitController *benefitController.BenefitHandler) {
	router.HandleFunc("/api/v0/benefit", benefitController.RegisterBenefit).Methods("POST")
	router.HandleFunc("/api/v0/benefit", benefitController.DeleteBenefit).Methods("DELETE")
	router.HandleFunc("/api/v0/benefit", benefitController.UpdateBenefit).Methods("PUT")
}
