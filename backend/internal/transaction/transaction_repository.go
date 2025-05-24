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

func (r *transactionRepository) Delete(id string) (*domain.Transaction, error) {
	query := `
		DELETE FROM transactions
		WHERE id = $1
		RETURNING id, user_id, asset_symbol, asset_type, quantity, price_in_usd, usd_brl_rate, price_in_brl, total_cost_usd, total_cost_brl, type, operation_date, created_at
	`

	var response domain.Transaction
	err := r.db.QueryRow(query, id).Scan(
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

func (r *transactionRepository) FindByID(id string) (*domain.Transaction, error) {
	query := `
		SELECT id, user_id, asset_symbol, asset_type, quantity, price_in_usd, usd_brl_rate, price_in_brl, total_cost_usd, total_cost_brl, type, operation_date, created_at
		FROM transactions
		WHERE id = $1
	`

	var response domain.Transaction
	err := r.db.QueryRow(query, id).Scan(
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

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (r *transactionRepository) FindAllBySymbol(userId, symbol string) ([]domain.Transaction, error) {
	query := `
		SELECT id, user_id, asset_symbol, asset_type, quantity, price_in_usd, usd_brl_rate, price_in_brl, total_cost_usd, total_cost_brl, type, operation_date, created_at
		FROM transactions
		WHERE user_id = $1 AND asset_symbol = $2
		ORDER BY operation_date ASC
	`
	var response []domain.Transaction
	rows, err := r.db.Query(query, userId, symbol)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction domain.Transaction
		err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.AssetSymbol, &transaction.AssetType, &transaction.Quantity, &transaction.PriceInUSD, &transaction.USDBRLRate, &transaction.PriceInBRL, &transaction.TotalCostUSD, &transaction.TotalCostBRL, &transaction.Type, &transaction.OperationDate, &transaction.CreatedAt)
		if err != nil {
			return nil, err
		}
		response = append(response, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return response, nil
}
