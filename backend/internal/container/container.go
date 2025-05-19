package container

import (
	"database/sql"

	"github.com/raphaelleveque/IRGlobal/backend/internal/user"
)

// Container mantém todas as dependências da aplicação
type Container struct {
	UserHandler *user.UserHandler
}

// NewContainer cria um novo container com todas as dependências configuradas
func NewContainer(db *sql.DB) *Container {
	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	return &Container{
		UserHandler: userHandler,
	}
}
