package store_handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/orbit-alliance/orbit-backend/internal/application/store/services"
)

type StoreHandler struct {
	registerService *services.RegisterStoreService
}

// Shared
func NewStoreHandler(registerService *services.RegisterStoreService) *StoreHandler {
	return &StoreHandler{
		registerService: registerService,
	}
}

// Shared
type RequestPayload struct {
	WalletAddress string `json:"wallet_address"`
	UserId        string `json:"user_id"`
}

func (h *StoreHandler) RegisterStore(w http.ResponseWriter, r *http.Request) {

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

	// code42 := parts[1]

	storeInfo, err := h.registerService.RegisterStore(r.Context(), payload.UserId, payload.WalletAddress)

	// fmt.Println(storeInfo)

	if err != nil {
		http.Error(w, "Erro ao registrar loja: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storeInfo)

	//Stopped here
}
