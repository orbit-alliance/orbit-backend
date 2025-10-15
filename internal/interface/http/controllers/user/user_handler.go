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
	IntraToken    string `json:"intra_token"`
}

func (h *UserHandler) Register42User(w http.ResponseWriter, r *http.Request) {


	var payload RequestPayload
	fmt.Println("Body:", r.Body)
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	code42 := payload.IntraToken

	basicInfo, err := h.registerService.Register42User(r.Context(), code42, payload.WalletAddress)
	if err != nil {
		http.Error(w, "Erro ao registrar usuário: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(basicInfo)
}
