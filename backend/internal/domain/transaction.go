package domain

import (
	"time"
)

type AssetType string
type OperationType string

const (
	Crypto AssetType = "CRYPTO"
	Stock  AssetType = "STOCK"
	ETF    AssetType = "ETF"
)

const (
	Buy  OperationType = "BUY"
	Sell OperationType = "SELL"
)

type Transaction struct {
	ID            string        `json:"id"`             // UUID
	UserID        string        `json:"user_id"`        // User UUID
	AssetSymbol   string        `json:"asset_symbol"`   // Asset symbol
	AssetType     AssetType     `json:"asset_type"`     // Asset type
	Quantity      float64       `json:"quantity"`       // Quantity
	PriceInUSD    float64       `json:"price_in_usd"`   // Price in USD
	USDBRLRate    float64       `json:"usd_brl_rate"`   // USD/BRL exchange rate
	PriceInBRL    float64       `json:"price_in_brl"`   // Price in BRL
	TotalCostUSD  float64       `json:"total_cost_usd"` // Total cost in USD
	TotalCostBRL  float64       `json:"total_cost_brl"` // Total cost in BRL
	Type          OperationType `json:"type"`           // Operation type
	OperationDate time.Time     `json:"operation_date"` // Operation date
	CreatedAt     time.Time     `json:"created_at"`     // Creation date
}

type TransactionService interface {
	AddTransaction(transaction *Transaction, dbTx DBTx) (*Transaction, error)
	DeleteTransaction(id string, dbTx DBTx) (*Transaction, error)
	FindByID(id string) (*Transaction, error)
	FindAllBySymbol(userId, symbol string) ([]Transaction, error)
	FindAllBySymbolExcludingOne(userId, symbol, transactionId string) ([]Transaction, error)
}

type TransactionRepository interface {
	Create(transaction *Transaction, dbTx DBTx) (*Transaction, error)
	Delete(id string, dbTx DBTx) (*Transaction, error)
	FindByID(id string) (*Transaction, error)
	FindAllBySymbol(userId, symbol string) ([]Transaction, error)
	FindAllBySymbolExcludingOne(userId, symbol, transactionId string) ([]Transaction, error)
}
