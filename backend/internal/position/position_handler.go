package position

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
)

type PositionHandler struct {
	positionService domain.PositionService
}

func NewPositionHandler(positionService domain.PositionService) *PositionHandler {
	return &PositionHandler{positionService: positionService}
}

// Register godoc
// @Summary      Lista as posições do usuário
// @Description  Retorna uma lista de posições para um usuário específico.
// @Tags         position
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Token de autenticação"
// @Success      200  {array}   domain.Position  "Lista de posições do usuário"
// @Failure      400  {object}  map[string]string "Dados inválidos"
// @Failure      500  {object}  map[string]string "Erro interno do servidor"
// @Router       /position/get [get]
// @Security     ApiKeyAuth
func (h *PositionHandler) GetPositions(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	positions, err := h.positionService.GetPositions(user.(*domain.User).ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"positions": positions,
	})
}
