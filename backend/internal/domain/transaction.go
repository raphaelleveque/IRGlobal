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
	UserID        string        `json:"user_id"`        // UUID do usuário
	AssetSymbol   string        `json:"asset_symbol"`   // Símbolo do ativo
	AssetType     AssetType     `json:"asset_type"`     // Tipo de ativo
	Quantity      float64       `json:"quantity"`       // Quantidade
	PriceInUSD    float64       `json:"price_in_usd"`   // Preço em USD
	USDBRLRate    float64       `json:"usd_brl_rate"`   // Taxa de câmbio USD/BRL
	PriceInBRL    float64       `json:"price_in_brl"`   // Preço em BRL
	TotalCostUSD  float64       `json:"total_cost_usd"` // Custo total em USD
	TotalCostBRL  float64       `json:"total_cost_brl"` // Custo total em USD
	Type          OperationType `json:"type"`           // Tipo de operação
	OperationDate time.Time     `json:"operation_date"` // Data da operação
	CreatedAt     time.Time     `json:"created_at"`     // Data de criação
}

type TransactionService interface {
	AddTransaction(transaction *Transaction) (*Transaction, error)
}

type TransactionRepository interface {
	Create(transaction *Transaction) (*Transaction, error)
}
