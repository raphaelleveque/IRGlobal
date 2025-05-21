package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raphaelleveque/IRGlobal/backend/internal/user"
)

func SetupRoutes(userHandler *user.UserHandler) *gin.Engine {
	router := gin.Default()

	setupAuthRoutes(router, userHandler)
	return router
}

func setupAuthRoutes(router *gin.Engine, userHandler *user.UserHandler) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", userHandler.Register)
	}
}
