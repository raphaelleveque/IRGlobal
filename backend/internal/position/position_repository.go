package position

import (
	"database/sql"

	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
)

// Interface para abstrair *sql.DB e *sql.Tx
type dbExecutor interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type positionRepository struct {
	db dbExecutor
}

func NewPositionRepository(db *sql.DB) domain.PositionRepository {
	return &positionRepository{db: db}
}

func NewPositionRepositoryWithTx(tx *sql.Tx) domain.PositionRepository {
	return &positionRepository{db: tx}
}

func (r *positionRepository) GetPositionByAssetSymbol(user_id, symbol string) (*domain.Position, error) {
	query := `
		SELECT id, user_id, asset_symbol, asset_type, quantity, average_cost_usd, average_cost_brl, total_cost_usd, total_cost_brl, created_at
		FROM positions
		WHERE user_id = $1 AND asset_symbol = $2
	`

	var position domain.Position
	err := r.db.QueryRow(query, user_id, symbol).Scan(
		&position.ID,
		&position.UserID,
		&position.AssetSymbol,
		&position.AssetType,
		&position.Quantity,
		&position.AverageCostUSD,
		&position.AverageCostBRL,
		&position.TotalCostUSD,
		&position.TotalCostBRL,
		&position.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &position, nil
}

func (r *positionRepository) UpdatePosition(position *domain.Position) (*domain.Position, error) {
	query := `
		INSERT INTO positions (user_id, asset_symbol, asset_type, quantity, average_cost_usd, average_cost_brl, total_cost_usd, total_cost_brl)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, user_id, asset_symbol, asset_type, quantity, average_cost_usd, average_cost_brl, total_cost_usd, total_cost_brl, created_at
	`

	err := r.db.QueryRow(query, position.UserID, position.AssetSymbol, position.AssetType, position.Quantity, position.AverageCostUSD, position.AverageCostBRL, position.TotalCostUSD, position.TotalCostBRL).Scan(
		&position.ID,
		&position.UserID,
		&position.AssetSymbol,
		&position.AssetType,
		&position.Quantity,
		&position.AverageCostUSD,
		&position.AverageCostBRL,
		&position.TotalCostUSD,
		&position.TotalCostBRL,
		&position.CreatedAt,
	)

	return position, err
}

func (r *positionRepository) DeletePosition(userId, symbol string) (*domain.Position, error) {
	query := `
		DELETE FROM positions
		WHERE user_id = $1 AND asset_symbol = $2
		RETURNING id, user_id, asset_symbol, asset_type, quantity, average_cost_usd, average_cost_brl, total_cost_usd, total_cost_brl, created_at
	`

	var position domain.Position
	err := r.db.QueryRow(query, userId, symbol).Scan(
		&position.ID,
		&position.UserID,
		&position.AssetSymbol,
		&position.AssetType,
		&position.Quantity,
		&position.AverageCostUSD,
		&position.AverageCostBRL,
		&position.TotalCostUSD,
		&position.TotalCostBRL,
		&position.CreatedAt,
	)

	return &position, err
}
