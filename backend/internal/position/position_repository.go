package position

import (
	"database/sql"

	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
	"github.com/raphaelleveque/IRGlobal/backend/internal/infrastructure"
)

type positionRepository struct {
	db *sql.DB
}

func NewPositionRepository(db *sql.DB) domain.PositionRepository {
	return &positionRepository{db: db}
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

func (r *positionRepository) GetPositions(user_id string) ([]domain.Position, error) {
	query := `
		SELECT id, user_id, asset_symbol, asset_type, quantity, average_cost_usd, average_cost_brl, total_cost_usd, total_cost_brl, created_at
		FROM positions
		WHERE user_id = $1
	`

	var positions []domain.Position
	rows, err := r.db.Query(query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var position domain.Position
		err := rows.Scan(
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
		if err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return positions, nil
}

func (r *positionRepository) UpdatePosition(position *domain.Position, dbTx domain.DBTx) (*domain.Position, error) {
	// UPSERT: UPDATE if exists, INSERT if not
	query := `
		INSERT INTO positions (user_id, asset_symbol, asset_type, quantity, average_cost_usd, average_cost_brl, total_cost_usd, total_cost_brl)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (user_id, asset_symbol)
		DO UPDATE SET
			quantity = EXCLUDED.quantity,
			average_cost_usd = EXCLUDED.average_cost_usd,
			average_cost_brl = EXCLUDED.average_cost_brl,
			total_cost_usd = EXCLUDED.total_cost_usd,
			total_cost_brl = EXCLUDED.total_cost_brl
		RETURNING id, user_id, asset_symbol, asset_type, quantity, average_cost_usd, average_cost_brl, total_cost_usd, total_cost_brl, created_at
	`

	// Convert DBTx to *sql.Tx
	sqlTx := dbTx.(*infrastructure.SqlDBTx).Tx

	err := sqlTx.QueryRow(query, position.UserID, position.AssetSymbol, position.AssetType, position.Quantity, position.AverageCostUSD, position.AverageCostBRL, position.TotalCostUSD, position.TotalCostBRL).Scan(
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

func (r *positionRepository) DeletePosition(userId, symbol string, dbTx domain.DBTx) (*domain.Position, error) {
	query := `
		DELETE FROM positions
		WHERE user_id = $1 AND asset_symbol = $2
		RETURNING id, user_id, asset_symbol, asset_type, quantity, average_cost_usd, average_cost_brl, total_cost_usd, total_cost_brl, created_at
	`

	// Convert DBTx to *sql.Tx
	sqlTx := dbTx.(*infrastructure.SqlDBTx).Tx

	var position domain.Position
	err := sqlTx.QueryRow(query, userId, symbol).Scan(
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
