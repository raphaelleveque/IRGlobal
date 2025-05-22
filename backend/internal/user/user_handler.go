package user

import (
	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
)

// UserHandler gerencia operações relacionadas ao usuário após autenticação
type UserHandler struct {
	userService domain.UserService
}

func NewUserHandler(userService domain.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// Funcionalidades CRUD de usuário serão adicionadas posteriormente
