package transaction

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
)

type TransactionHandler struct {
	transactionService domain.TransactionService
}

type AddTransactionRequest struct {
	AssetSymbol   string               `json:"asset_symbol" binding:"required" example:"AAPL"`       // Símbolo do ativo
	AssetType     domain.AssetType     `json:"asset_type" binding:"required,oneof=CRYPTO STOCK ETF"` // Tipo de ativo
	Quantity      float64              `json:"quantity" binding:"required,min=0"`                    // Quantidade
	PriceInUSD    float64              `json:"price_in_usd" binding:"required,min=0"`                // Preço em USD
	Type          domain.OperationType `json:"type" binding:"required,oneof=BUY SELL"`               // Tipo de operação
	OperationDate string               `json:"operation_date" binding:"required"`                    // Data da operação
}

func NewTransactionHandler(transactionService domain.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService: transactionService}
}

// Register godoc
// @Summary      Adiciona uma nova Transação
// @Description  Adiciona uma nova Transação ao sistema
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Token de autenticação"
// @Param        request body AddTransactionRequest true "Dados da transação"
// @Success      201  {object}  domain.Transaction  "Transação criada com sucesso"
// @Failure      400  {object}  map[string]string    "Dados inválidos"
// @Failure      500  {object}  map[string]string    "Erro interno do servidor"
// @Router       /transaction/add [post]
// @Security     ApiKeyAuth
func (h *TransactionHandler) AddTransaction(c *gin.Context) {
	var req AddTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	operationDate, err := time.Parse("2006-01-02", req.OperationDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	transaction := &domain.Transaction{
		UserID:        user.(*domain.User).ID,
		AssetSymbol:   req.AssetSymbol,
		AssetType:     req.AssetType,
		Quantity:      req.Quantity,
		PriceInUSD:    req.PriceInUSD,
		Type:          req.Type,
		OperationDate: operationDate,
	}

	transaction, err = h.transactionService.AddTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"transaction": transaction,
	})
}
