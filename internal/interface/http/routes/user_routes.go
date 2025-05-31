package routes

import (
	"github.com/gorilla/mux"
	usercontroller "github.com/orbit-alliance/orbit-backend/internal/interface/http/controllers/user"
)

// RegisterUserRoutes registra as rotas relacionadas ao usuário no roteador fornecido
// e associa o Auth42Handler para lidar com as requisições de autenticação via 42.
// Agora a função recebe o handler como parâmetro
func RegisterUserRoutes(router *mux.Router, auth42Handler *usercontroller.Auth42Handler) {
	router.HandleFunc("/auth/42/login", auth42Handler.Login).Methods("GET")
	router.HandleFunc("/auth/42/callback", auth42Handler.Callback).Methods("GET")
}
