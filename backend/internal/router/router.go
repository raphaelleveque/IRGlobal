package router

import (
	"github.com/gorilla/mux"
	"github.com/raphaelleveque/IRGlobal/backend/internal/user"
)

// SetupRoutes configura todas as rotas da aplicação
func SetupRoutes(userHandler *user.UserHandler) *mux.Router {
	router := mux.NewRouter()

	// Configurar rotas de autenticação
	setupAuthRoutes(router, userHandler)

	// Configurar rotas protegidas de usuários
	setupProtectedUserRoutes(router, userHandler)

	return router
}

// setupAuthRoutes configura as rotas de autenticação
func setupAuthRoutes(router *mux.Router, userHandler *user.UserHandler) {
	router.HandleFunc("/auth/register", userHandler.Register).Methods("POST")
	router.HandleFunc("/auth/login", userHandler.Login).Methods("POST")
}

// setupProtectedUserRoutes configura as rotas protegidas de usuários
func setupProtectedUserRoutes(router *mux.Router, userHandler *user.UserHandler) {
	router.HandleFunc("/users", userHandler.AuthMiddleware(userHandler.GetAllUsers)).Methods("GET")
}
