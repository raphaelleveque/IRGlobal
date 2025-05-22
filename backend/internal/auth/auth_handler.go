package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
)

type AuthHandler struct {
	userService domain.UserService
	authService domain.AuthService
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required" example:"João Silva"`
	Email    string `json:"email" binding:"required,email" example:"joao@example.com"`
	Password string `json:"password" binding:"required,min=6" example:"senha123"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"joao@example.com"`
	Password string `json:"password" binding:"required,min=6" example:"senha123"`
}

func NewAuthHandler(userService domain.UserService, authService domain.AuthService) *AuthHandler {
	return &AuthHandler{userService: userService, authService: authService}
}

// GetAuthService retorna o serviço de autenticação
func (h *AuthHandler) GetAuthService() domain.AuthService {
	return h.authService
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
func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	createdUser, err := h.userService.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authService.GenerateToken(createdUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao gerar token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user":  createdUser,
		"token": token,
	})
}

// Login godoc
// @Summary      Autentica o usuário
// @Description  Autentica o usuário no sistema
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body LoginRequest true "Dados do usuário"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.GetByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	isValidPassword := h.authService.ValidatePassword(req.Password, user.Password)
	if !isValidPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Senha inválida"})
		return
	}

	token, err := h.authService.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}
