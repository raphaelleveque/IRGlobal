package domain

import "time"

type Position struct {
	ID             string    `json:"id"`               // UUID
	UserID         string    `json:"user_id"`          // UUID do usuário
	AssetSymbol    string    `json:"asset_symbol"`     // Símbolo do ativo
	AssetType      AssetType `json:"asset_type"`       // Tipo de ativo
	Quantity       float64   `json:"quantity"`         // Quantidade
	AverageCostUSD float64   `json:"average_cost_usd"` // Preço em BRL
	AverageCostBRL float64   `json:"average_cost_brl"` // Preço em BRL
	TotalCostUSD   float64   `json:"total_cost_usd"`   // Custo total em USD
	TotalCostBRL   float64   `json:"total_cost_brl"`   // Custo total em USD
	CreatedAt      time.Time `json:"created_at"`       // Data de criação
}

type PositionService interface {
	CalculatePosition(transaction *Transaction, dbTx DBTx) (*Position, error)
	RecalculatePosition(userId, symbol, transactionId string, dbTx DBTx) (*Position, error)
	GetPositionByAssetSymbol(userId, symbol string) (*Position, error)
}

type PositionRepository interface {
	UpdatePosition(position *Position, dbTx DBTx) (*Position, error)
	GetPositionByAssetSymbol(user_id, symbol string) (*Position, error)
	DeletePosition(userId, symbol string, dbTx DBTx) (*Position, error)
}
