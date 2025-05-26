package user

import (
	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
)

// UserHandler manages user-related operations after authentication
type UserHandler struct {
	userService domain.UserService
}

func NewUserHandler(userService domain.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// User CRUD functionalities will be added later
