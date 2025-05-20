package container

import (
	"database/sql"

	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
	"github.com/raphaelleveque/IRGlobal/backend/internal/user"
)

type AppContainer struct {
	userService domain.UserService
	userHandler *user.UserHandler
}

func NewAppContainer(db *sql.DB) *AppContainer {
	repo := user.NewUserRepository(db)
	userService := user.NewUserService(repo)
	userHandler := user.NewUserHandler(userService)

	return &AppContainer{
		userService: userService,
		userHandler: userHandler,
	}
}

func (c *AppContainer) GetUserHandler() *user.UserHandler {
	return c.userHandler
}
