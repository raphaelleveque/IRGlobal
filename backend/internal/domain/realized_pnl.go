package domain

import "time"

type RealizedPNL struct {
	ID                string    `json:"id"`                   // UUID
	UserID            string    `json:"user_id"`              // UUID do usuário
	AssetSymbol       string    `json:"asset_symbol"`         // Símbolo do ativo
	AssetType         AssetType `json:"asset_type"`           // Tipo de ativo
	Quantity          float64   `json:"quantity"`             // Quantidade
	AverageCostUSD    float64   `json:"average_cost_usd"`     // Custo médio em USD
	AverageCostBRL    float64   `json:"average_cost_brl"`     // Custo médio em BRL
	TotalCostUSD      float64   `json:"total_cost_usd"`       // Custo total em USD
	TotalCostBRL      float64   `json:"total_cost_brl"`       // Custo total em BRL
	SellingPriceUSD   float64   `json:"selling_price_usd"`    // Preço de venda em USD
	SellingPriceBRL   float64   `json:"selling_price_brl"`    // Preço de venda em BRL
	TotalValueSoldUSD float64   `json:"total_value_sold_usd"` // Valor total da venda em USD
	TotalValueSoldBRL float64   `json:"total_value_sold_brl"` // Valor total da venda em BRL
	RealizedProfitUSD float64   `json:"realized_profit_usd"`  // Lucro realizado em USD
	RealizedProfitBRL float64   `json:"realized_profit_brl"`  // Lucro realizado em BRL
	CreatedAt         time.Time `json:"created_at"`           // Data de criação
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
