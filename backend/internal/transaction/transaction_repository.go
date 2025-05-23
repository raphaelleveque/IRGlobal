package transaction

import (
	"database/sql"

	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
)

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) domain.TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(transaction *domain.Transaction) (*domain.Transaction, error) {
	query := `
		INSERT INTO transactions (user_id, asset_symbol, asset_type, quantity, price_in_usd, usd_brl_rate, price_in_brl, total_cost_usd, total_cost_brl, type, operation_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id, user_id, asset_symbol, asset_type, quantity, price_in_usd, usd_brl_rate, price_in_brl, total_cost_usd, total_cost_brl, type, operation_date, created_at
		`
	var response domain.Transaction
	err := r.db.QueryRow(query, transaction.UserID, transaction.AssetSymbol, transaction.AssetType, transaction.Quantity, transaction.PriceInUSD, transaction.USDBRLRate, transaction.PriceInBRL, transaction.TotalCostUSD, transaction.TotalCostBRL, transaction.Type, transaction.OperationDate).Scan(
		&response.ID,
		&response.UserID,
		&response.AssetSymbol,
		&response.AssetType,
		&response.Quantity,
		&response.PriceInUSD,
		&response.USDBRLRate,
		&response.PriceInBRL,
		&response.TotalCostUSD,
		&response.TotalCostBRL,
		&response.Type,
		&response.OperationDate,
		&response.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}