package domain

import "time"

type Position struct {
	ID             string    `json:"id"`               // UUID
	UserID         string    `json:"user_id"`          // User UUID
	AssetSymbol    string    `json:"asset_symbol"`     // Asset symbol
	AssetType      AssetType `json:"asset_type"`       // Asset type
	Quantity       float64   `json:"quantity"`         // Quantity
	AverageCostUSD float64   `json:"average_cost_usd"` // Average cost in USD
	AverageCostBRL float64   `json:"average_cost_brl"` // Average cost in BRL
	TotalCostUSD   float64   `json:"total_cost_usd"`   // Total cost in USD
	TotalCostBRL   float64   `json:"total_cost_brl"`   // Total cost in BRL
	CreatedAt      time.Time `json:"created_at"`       // Creation date
}

type PositionService interface {
	CalculatePosition(transaction *Transaction, dbTx DBTx) (*Position, error)
	RecalculatePosition(userId, symbol, transactionId string, dbTx DBTx) (*Position, error)
	GetPositionByAssetSymbol(userId, symbol string) (*Position, error)
	GetPositions(userId string) ([]Position, error)
}

type PositionRepository interface {
	UpdatePosition(position *Position, dbTx DBTx) (*Position, error)
	GetPositionByAssetSymbol(user_id, symbol string) (*Position, error)
	GetPositions(user_id string) ([]Position, error)
	DeletePosition(userId, symbol string, dbTx DBTx) (*Position, error)
}
