package user_controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/orbit-alliance/orbit-backend/internal/application/user/services"
)

type UserHandler struct {
	registerService *services.Register42Service
}

func NewUserHandler(registerService *services.Register42Service) *UserHandler {
	return &UserHandler{
		registerService: registerService,
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

	id42, login, err := h.registerService.Register42User(r.Context(), code42, payload.WalletAddress)
	if err != nil {
		http.Error(w, "Erro ao registrar usuário: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"id42":  id42,
		"login": login,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
