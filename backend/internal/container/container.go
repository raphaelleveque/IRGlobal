package container

import (
	"database/sql"

	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
	"github.com/raphaelleveque/IRGlobal/backend/internal/user"
)

// AppContainer mantém todas as dependências da aplicação
type AppContainer struct {
	userService domain.UserService
	userHandler *user.UserHandler
}

// NewAppContainer cria um novo container com todas as dependências configuradas
func NewAppContainer(db *sql.DB) *AppContainer {
	// Inicializa as dependências em ordem: repository -> service -> handler
	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	return &AppContainer{
		userService: userService,
		userHandler: userHandler,
	}
}

// GetUserHandler retorna o handler de usuários configurado
func (c *AppContainer) GetUserHandler() *user.UserHandler {
	return c.userHandler
}
