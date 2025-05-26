package transaction

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
	"github.com/raphaelleveque/IRGlobal/backend/internal/transaction/orchestrator"
)

type TransactionHandler struct {
	transactionService domain.TransactionService
	orchestrator       *orchestrator.TransactionOrchestrator
}

type AddTransactionRequest struct {
	AssetSymbol   string               `json:"asset_symbol" binding:"required" example:"AAPL"`                       // Asset symbol
	AssetType     domain.AssetType     `json:"asset_type" binding:"required,oneof=CRYPTO STOCK ETF" example:"STOCK"` // Asset type
	Quantity      float64              `json:"quantity" binding:"required,min=0" example:"20"`                       // Quantity
	PriceInUSD    float64              `json:"price_in_usd" binding:"required,min=0" example:"50"`                   // Price in USD
	Type          domain.OperationType `json:"type" binding:"required,oneof=BUY SELL" example:"BUY"`                 // Operation type
	OperationDate string               `json:"operation_date" binding:"required" example:"2025-02-11"`               // Operation date
}

type DeleteTransactionRequest struct {
	ID string `json:"id" binding:"required" example:"d081b7c0-b3b6-49ba-a9b7-86b56a65fb89"`
}

func NewTransactionHandler(transactionService domain.TransactionService, orchestrator *orchestrator.TransactionOrchestrator) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
		orchestrator:       orchestrator,
	}
}

// Register godoc
// @Summary      Adiciona uma nova Transação
// @Description  Adiciona uma nova Transação ao sistema
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Authentication token"
// @Param        request body AddTransactionRequest true "Transaction data"
// @Success      201  {object}  domain.Transaction  "Transaction created successfully"
// @Failure      400  {object}  map[string]string    "Invalid data"
// @Failure      500  {object}  map[string]string    "Internal server error"
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

	transaction, position, realizedPnl, err := h.orchestrator.AddTransactionWithPosition(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"transaction": transaction,
		"position":    position,
		"realizedPnl": realizedPnl,
	})
}

// Register godoc
// @Summary      Deletes a Transaction
// @Description  Deletes a Transaction from the system
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Authentication Token"
// @Param        request body DeleteTransactionRequest true "Transaction details"
// @Success      200  {object}  domain.Transaction  "Transaction successfully deleted"
// @Failure      404  {object}  map[string]string    "Transaction not found"
// @Failure      500  {object}  map[string]string    "Internal server error"
// @Router       /transaction/delete [delete]
// @Security     ApiKeyAuth
func (h *TransactionHandler) DeleteTransaction(c *gin.Context) {
	var req DeleteTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if _, err := h.transactionService.FindByID(req.ID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	transaction, position, pnl, err := h.orchestrator.DeleteTransactionWithPosition(req.ID, user.(*domain.User).ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transaction": transaction,
		"position":    position,
		"pnl":         pnl,
	})
}


// Register godoc
// @Summary      List user transactions
// @Description  Returns a list of transactions for a specific user.
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Authentication token"
// @Success      200  {array}   domain.Position  "List of user positions"
// @Failure      400  {object}  map[string]string "Invalid data"
// @Failure      500  {object}  map[string]string "Internal server error"
// @Router       /transaction/get [get]
// @Security     ApiKeyAuth
func (h *TransactionHandler) GetTransactions(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	transactions, err := h.transactionService.FindAll(user.(*domain.User).ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
	})
}

