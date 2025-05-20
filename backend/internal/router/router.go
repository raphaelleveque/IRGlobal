package router

import (
	"github.com/gorilla/mux"
	"github.com/raphaelleveque/IRGlobal/backend/internal/user"
)

func SetupRoutes(userHandler *user.UserHandler) *mux.Router  {
	router := mux.NewRouter()

	setupAuthRoutes(router, userHandler)
	return router
}

func setupAuthRoutes(router *mux.Router, userHandler *user.UserHandler) {
	router.HandleFunc("/auth/register", userHandler.Register).Methods("POST")
}