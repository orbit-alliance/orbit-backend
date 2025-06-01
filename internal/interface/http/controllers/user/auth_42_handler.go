package user_controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
)

type Auth42Handler struct {
	auth42Gateway user.Api42Gateway
}

func NewAuth42Handler(auth42Gateway user.Api42Gateway) *Auth42Handler {
	return &Auth42Handler{
		auth42Gateway: auth42Gateway,
	}
}

// ➕ Redireciona para o login da 42
func (h *Auth42Handler) Login(w http.ResponseWriter, r *http.Request) {
	clientID := os.Getenv("CLIENT_ID_42")
	redirectURI := os.Getenv("REDIRECT_URI_42")

	if clientID == "" || redirectURI == "" {
		log.Fatal("CLIENT_ID_42 ou REDIRECT_URI_42 não estão definidos")
	}

	authURL := fmt.Sprintf(
		"https://api.intra.42.fr/oauth/authorize?client_id=%s&redirect_uri=%s&response_type=code",
		clientID,
		url.QueryEscape(redirectURI),
	)

	fmt.Println("Auth URL gerada:", authURL)

	http.Redirect(w, r, authURL, http.StatusFound)
}

// 🔥 Callback que recebe o code, troca pelo token e busca o user
func (h *Auth42Handler) Callback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Código de autorização não encontrado", http.StatusBadRequest)
		return
	}

	// 🪄 🔗 Usa o gateway para trocar o code por token
	token, err := h.auth42Gateway.ExchangeCodeForToken(r.Context(), code)
	if err != nil {
		http.Error(w, "Erro ao obter token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 🔍 Consulta o user na API da 42
	id42, login, err := h.auth42Gateway.GetBasicUserInfo(r.Context(), token)
	if err != nil {
		http.Error(w, "Erro ao obter dados do usuário: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 🟢 Retorna ID e Login
	response := map[string]string{
		"id42":  id42,
		"login": login,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
