package user_controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/orbit-alliance/orbit-backend/internal/application/user/services"
	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
)

type UserHandler struct {
	registerService			*services.Register42Service
	transferCoinsService	*services.TransferCoinsService
	buyBenefitService		*services.BuyBenefitService
}

func NewUserHandler(
	registerService *services.Register42Service, 
	transferCoinsService *services.TransferCoinsService,  
	buyBenefitService *services.BuyBenefitService) *UserHandler {
	return &UserHandler{
		registerService:		registerService,
		transferCoinsService:	transferCoinsService,
		buyBenefitService:		buyBenefitService,
	}
}

type RequestPayload struct {
	WalletAddress string `json:"wallet_address"`
}

func (h *UserHandler) Register42User(w http.ResponseWriter, r *http.Request) {

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		http.Error(w, "Authorization header format must be Bearer {token}", http.StatusUnauthorized)
		return
	}

	var payload RequestPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	code42 := parts[1]

	basicInfo, err := h.registerService.Register42User(r.Context(), code42, payload.WalletAddress)
	if err != nil {
		http.Error(w, "Erro ao registrar usuário: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(basicInfo)
}

type TransferCoinsPayload struct {
	From		string	`json:"from"`
	To			string	`json:"to"`
	Amount		uint64	`json:"amount"`
}


func (h *UserHandler) TransferCoins(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		http.Error(w, "Authorization header format must be Bearer {token}", http.StatusUnauthorized)
		return
	}

	var     payload TransferCoinsPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	transferDTO := user.NewTransferDTO(payload.From, payload.To, payload.Amount)
	err = h.transferCoinsService.TransferCoins(*transferDTO)
	if err != nil {
		http.Error(w, "Could not transfer coins", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transferDTO)
}

type BuyBenfitPayload struct {
	UserId		string	`json:"user_id"`
	BenefitId	string	`json:"benefit_id"`
}

func (h *UserHandler) UserBuyBenefit(w http.ResponseWriter, r *http.Request) {
	
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		http.Error(w, "Authorization header format must be Bearer {token}", http.StatusUnauthorized)
		return
	}

	var payload BuyBenfitPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	purchaseDT0, err := h.buyBenefitService.BuyBenefit(payload.UserId, payload.BenefitId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(purchaseDT0)
}
