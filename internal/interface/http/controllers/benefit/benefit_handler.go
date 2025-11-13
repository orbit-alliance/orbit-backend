package benefit_handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/orbit-alliance/orbit-backend/internal/application/benefit/services"
	"github.com/orbit-alliance/orbit-backend/internal/domain/benefit"
)

type BenefitHandler struct {
	registerService *services.RegisterBenefitService
} 

// Shared
func NewBenefitHandler(registerService *services.RegisterBenefitService) *BenefitHandler {
	return &BenefitHandler{
		registerService: registerService,
	}
}

func CheckAuthorizarionHeader(w http.ResponseWriter, r *http.Request) bool {
	
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return false
	}
	
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		http.Error(w, "Authorization header format must be Bearer {token}", http.StatusUnauthorized)
		return false
	}
	return true
}

func (h *BenefitHandler) RegisterBenefit(w http.ResponseWriter, r *http.Request) {

	if (!CheckAuthorizarionHeader(w, r)) {
		return
	}

	var payload benefit.BenefitPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	
	benefitInfo, err := h.registerService.RegisterBenefit(r.Context(), payload)
	if err != nil {
		http.Error(w, "Erro ao registrar beneficio: " + err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(benefitInfo)
}

func (h *BenefitHandler) UpdateBenefit(w http.ResponseWriter, r *http.Request) {

	if (!CheckAuthorizarionHeader(w, r)) {
		return
	}

	var payload benefit.BenefitPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	benefitInfo, err := h.registerService.UpdateBenefit(r.Context(), payload)
	if err != nil {
		http.Error(w, "Erro ao registrar beneficio: " + err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(benefitInfo)
}

func (h *BenefitHandler) DeleteBenefit(w http.ResponseWriter, r *http.Request) {

	if (!CheckAuthorizarionHeader(w, r)) {
		return
	}
	
	var payload benefit.BenefitPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	benefitDTO, err := h.registerService.DeleteBenefit(r.Context(), payload)
	if err != nil {
		http.Error(w, "Erro ao deletar beneficio: " + err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(benefitDTO)
}
