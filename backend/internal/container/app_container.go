package container

import (
	"database/sql"

	"github.com/raphaelleveque/IRGlobal/backend/internal/auth"
	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
	"github.com/raphaelleveque/IRGlobal/backend/internal/user"
)

type AppContainer struct {
	userService domain.UserService
	userHandler *user.UserHandler
	authService domain.AuthService
	authHandler *auth.AuthHandler
}

func NewAppContainer(db *sql.DB, secretKey []byte) *AppContainer {
	// Repositories
	userRepo := user.NewUserRepository(db)

	// Services
	userService := user.NewUserService(userRepo)
	authService := auth.NewAuthService(secretKey, userService)

	// Handlers
	userHandler := user.NewUserHandler(userService)
	authHandler := auth.NewAuthHandler(userService, authService)

	return &AppContainer{
		userService: userService,
		userHandler: userHandler,
		authService: authService,
		authHandler: authHandler,
	}
}

func (c *AppContainer) GetUserHandler() *user.UserHandler {
	return c.userHandler
}

func (c *AppContainer) GetAuthHandler() *auth.AuthHandler {
	return c.authHandler
}

func (c *AppContainer) GetAuthService() domain.AuthService {
	return c.authService
}
