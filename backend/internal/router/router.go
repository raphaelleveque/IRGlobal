package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raphaelleveque/IRGlobal/backend/internal/auth"
	"github.com/raphaelleveque/IRGlobal/backend/internal/container"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(appContainer *container.AppContainer) *gin.Engine {
	router := gin.Default()

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Public routes (do not require authentication)
	setupAuthRoutes(router, appContainer)

	// Rotas protegidas (necessitam de autenticação)
	setupProtectedRoutes(router, appContainer)

	return router
}

// setupAuthRoutes configura as rotas públicas de autenticação
func setupAuthRoutes(router *gin.Engine, appContainer *container.AppContainer) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", appContainer.GetAuthHandler().Register)
		auth.POST("/login", appContainer.GetAuthHandler().Login)
	}
}

// setupProtectedRoutes configura as rotas protegidas
func setupProtectedRoutes(router *gin.Engine, appContainer *container.AppContainer) {
	// Middleware de autenticação será aplicado a este grupo
	protected := router.Group("/")
	protected.Use(auth.AuthMiddleware(appContainer.GetAuthService()))

	// Protected transaction routes
	transaction := protected.Group("/transaction")
	{
		transaction.POST("/add", appContainer.GetTransactionHandler().AddTransaction)
		transaction.DELETE("/delete", appContainer.GetTransactionHandler().DeleteTransaction)
	}

	// Position routes
	position := protected.Group("/position")
	{
		position.GET("get", appContainer.GetPositionHandler().GetPositions)
	}

	pnl := protected.Group("/pnl")
	{
		pnl.GET("/get", appContainer.GetRealizedPNLHandler().GetPNL)
	}
}
