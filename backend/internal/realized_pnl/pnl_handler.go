package realizedpnl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
)

type RealizedPNLHandler struct {
	service domain.RealizedPNLService
}

func NewRealizedPNLHandler(service domain.RealizedPNLService) *RealizedPNLHandler {
	return &RealizedPNLHandler{service: service}
}

// Register godoc
// @Summary      List user Profit and Loss
// @Description  Returns a list of PNL for a specific user.
// @Tags         pnl
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Authentication token"
// @Success      200  {array}   domain.RealizedPNL  "List of user PNL"
// @Failure      400  {object}  map[string]string "Invalid data"
// @Failure      500  {object}  map[string]string "Internal server error"
// @Router       /realized-pnl/get [get]
// @Security     ApiKeyAuth
func (h *RealizedPNLHandler) GetPNL(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	pnls, err := h.service.GetPNLs(user.(*domain.User).ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"pnls": pnls,
	})
}
