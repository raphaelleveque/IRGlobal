package router

import (
	"github.com/gorilla/mux"
	"github.com/raphaelleveque/IRGlobal/backend/internal/user"
)

// SetupRoutes configura todas as rotas da aplicação
func SetupRoutes(userHandler *user.UserHandler) *mux.Router {
	router := mux.NewRouter()

	// Rotas públicas
	router.HandleFunc("/register", userHandler.Register).Methods("POST")
	router.HandleFunc("/login", userHandler.Login).Methods("POST")

	// Rotas protegidas
	router.HandleFunc("/users", userHandler.AuthMiddleware(userHandler.GetAllUsers)).Methods("GET")

	return router
}
