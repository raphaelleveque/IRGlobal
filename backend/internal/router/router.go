package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raphaelleveque/IRGlobal/backend/internal/user"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(userHandler *user.UserHandler) *gin.Engine {
	router := gin.Default()

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	setupAuthRoutes(router, userHandler)
	return router
}

func setupAuthRoutes(router *gin.Engine, userHandler *user.UserHandler) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", userHandler.Register)
	}
}
