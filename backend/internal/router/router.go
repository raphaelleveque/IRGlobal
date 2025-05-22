package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raphaelleveque/IRGlobal/backend/internal/auth"
	"github.com/raphaelleveque/IRGlobal/backend/internal/user"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(authHandler *auth.AuthHandler, userHandler *user.UserHandler) *gin.Engine {
	router := gin.Default()

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Rotas públicas (não necessitam de autenticação)
	setupAuthRoutes(router, authHandler)

	// Rotas protegidas (necessitam de autenticação)
	// setupProtectedRoutes(router, userHandler, authHandler.GetAuthService())

	return router
}

// setupAuthRoutes configura as rotas públicas de autenticação
func setupAuthRoutes(router *gin.Engine, authHandler *auth.AuthHandler) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		// Outras rotas de autenticação serão adicionadas posteriormente
		// Exemplo: auth.POST("/login", authHandler.Login)
	}
}
