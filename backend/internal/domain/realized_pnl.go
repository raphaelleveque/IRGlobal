package domain

import "time"

type RealizedPNL struct {
	ID                string    `json:"id"`                   // UUID
	UserID            string    `json:"user_id"`              // User UUID
	AssetSymbol       string    `json:"asset_symbol"`         // Asset symbol
	AssetType         AssetType `json:"asset_type"`           // Asset type
	Quantity          float64   `json:"quantity"`             // Quantity
	AverageCostUSD    float64   `json:"average_cost_usd"`     // Average cost in USD
	AverageCostBRL    float64   `json:"average_cost_brl"`     // Average cost in BRL
	TotalCostUSD      float64   `json:"total_cost_usd"`       // Total cost in USD
	TotalCostBRL      float64   `json:"total_cost_brl"`       // Total cost in BRL
	SellingPriceUSD   float64   `json:"selling_price_usd"`    // Selling price in USD
	SellingPriceBRL   float64   `json:"selling_price_brl"`    // Selling price in BRL
	TotalValueSoldUSD float64   `json:"total_value_sold_usd"` // Total value sold in USD
	TotalValueSoldBRL float64   `json:"total_value_sold_brl"` // Total value sold in BRL
	RealizedProfitUSD float64   `json:"realized_profit_usd"`  // Realized profit in USD
	RealizedProfitBRL float64   `json:"realized_profit_brl"`  // Realized profit in BRL
	CreatedAt         time.Time `json:"created_at"`           // Creation date
}

type RealizedPNLService interface {
	CalculatePNL(transaction *Transaction, position *Position, dbTx DBTx) (*RealizedPNL, error)
	RecalculatePNL(userId, symbol, transactionId string, dbTx DBTx) (*RealizedPNL, error)
	GetPNLs(userId string) ([]RealizedPNL, error)
}

type RealizedPNLRepository interface {
	UpdatePNL(pnl *RealizedPNL, dbTx DBTx) (*RealizedPNL, error)
	GetPNLByAssetSymbol(user_id, symbol string) (*RealizedPNL, error)
	GetPNLs(userId string) ([]RealizedPNL, error)
	DeletePNL(userId, symbol string, dbTx DBTx) (*RealizedPNL, error)
}
