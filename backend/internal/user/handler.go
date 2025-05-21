package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
)

type UserHandler struct {
	userService domain.UserService
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required" example:"João Silva"`
	Email    string `json:"email" binding:"required,email" example:"joao@example.com"`
	Password string `json:"password" binding:"required,min=6" example:"senha123"`
}

func NewUserHandler(userService domain.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// Register godoc
// @Summary      Registrar um novo usuário
// @Description  Registra um novo usuário no sistema
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body RegisterRequest true "Dados do usuário"
// @Success      201  {object}  domain.User
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /auth/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.userService.Register(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
