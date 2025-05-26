package realizedpnl

import (
	"database/sql"

	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
	"github.com/raphaelleveque/IRGlobal/backend/internal/infrastructure"
)

type realizedPNLRepository struct {
	db *sql.DB
}

func NewRealizedPNLRepository(db *sql.DB) domain.RealizedPNLRepository {
	return &realizedPNLRepository{db: db}
}

func (r *realizedPNLRepository) UpdatePNL(pnl *domain.RealizedPNL, dbTx domain.DBTx) (*domain.RealizedPNL, error) {
	// UPSERT: UPDATE if exists, INSERT if not
	query := `
		INSERT INTO realized_pnl (user_id, asset_symbol, asset_type, quantity, average_cost_usd, average_cost_brl, total_cost_usd, total_cost_brl, selling_price_usd, selling_price_brl, total_value_sold_usd, total_value_sold_brl, realized_profit_usd, realized_profit_brl)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
		ON CONFLICT (user_id, asset_symbol)
		DO UPDATE SET
			quantity = EXCLUDED.quantity,
			average_cost_usd = EXCLUDED.average_cost_usd,
			average_cost_brl = EXCLUDED.average_cost_brl,
			total_cost_usd = EXCLUDED.total_cost_usd,
			total_cost_brl = EXCLUDED.total_cost_brl,
			selling_price_usd = EXCLUDED.selling_price_usd,
			selling_price_brl = EXCLUDED.selling_price_brl,
			total_value_sold_usd = EXCLUDED.total_value_sold_usd,
			total_value_sold_brl = EXCLUDED.total_value_sold_brl,
			realized_profit_usd = EXCLUDED.realized_profit_usd,
			realized_profit_brl = EXCLUDED.realized_profit_brl
		RETURNING id, user_id, asset_symbol, asset_type, quantity, average_cost_usd, average_cost_brl, total_cost_usd, total_cost_brl, selling_price_usd, selling_price_brl, total_value_sold_usd, total_value_sold_brl, realized_profit_usd, realized_profit_brl, created_at
	`

	// Convert DBTx to *sql.Tx
	sqlTx := dbTx.(*infrastructure.SqlDBTx).Tx

	err := sqlTx.QueryRow(query, pnl.UserID, pnl.AssetSymbol, pnl.AssetType, pnl.Quantity, pnl.AverageCostUSD, pnl.AverageCostBRL, pnl.TotalCostUSD, pnl.TotalCostBRL, pnl.SellingPriceUSD, pnl.SellingPriceBRL, pnl.TotalValueSoldUSD, pnl.TotalValueSoldBRL, pnl.RealizedProfitUSD, pnl.RealizedProfitBRL).Scan(
		&pnl.ID,
		&pnl.UserID,
		&pnl.AssetSymbol,
		&pnl.AssetType,
		&pnl.Quantity,
		&pnl.AverageCostUSD,
		&pnl.AverageCostBRL,
		&pnl.TotalCostUSD,
		&pnl.TotalCostBRL,
		&pnl.SellingPriceUSD,
		&pnl.SellingPriceBRL,
		&pnl.TotalValueSoldUSD,
		&pnl.TotalValueSoldBRL,
		&pnl.RealizedProfitUSD,
		&pnl.RealizedProfitBRL,
		&pnl.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return pnl, nil

}

func (r *realizedPNLRepository) GetPNLByAssetSymbol(user_id, symbol string) (*domain.RealizedPNL, error) {
	query := `
		SELECT 
			id,
			user_id,
			asset_symbol,
			asset_type,
			quantity,
			average_cost_usd,
			average_cost_brl,
			total_cost_usd,
			total_cost_brl,
			selling_price_usd,
			selling_price_brl,
			total_value_sold_usd,
			total_value_sold_brl,
			realized_profit_usd,
			realized_profit_brl,
			created_at
		FROM realized_pnl
		WHERE user_id = $1 AND asset_symbol = $2
	`

	var pnl domain.RealizedPNL
	err := r.db.QueryRow(query, user_id, symbol).Scan(
		&pnl.ID,
		&pnl.UserID,
		&pnl.AssetSymbol,
		&pnl.AssetType,
		&pnl.Quantity,
		&pnl.AverageCostUSD,
		&pnl.AverageCostBRL,
		&pnl.TotalCostUSD,
		&pnl.TotalCostBRL,
		&pnl.SellingPriceUSD,
		&pnl.SellingPriceBRL,
		&pnl.TotalValueSoldUSD,
		&pnl.TotalValueSoldBRL,
		&pnl.RealizedProfitUSD,
		&pnl.RealizedProfitBRL,
		&pnl.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &pnl, nil
}

func (r *realizedPNLRepository) DeletePNL(userId, symbol string, dbTx domain.DBTx) (*domain.RealizedPNL, error) {
	query := `
		DELETE FROM realized_pnl
		WHERE user_id = $1 AND asset_symbol = $2
		RETURNING id, user_id, asset_symbol, asset_type, quantity, average_cost_usd, average_cost_brl, total_cost_usd, total_cost_brl, selling_price_usd, selling_price_brl, total_value_sold_usd, total_value_sold_brl, realized_profit_usd, realized_profit_brl, created_at
	`

	sqlTx := dbTx.(*infrastructure.SqlDBTx).Tx

	var pnl domain.RealizedPNL
	err := sqlTx.QueryRow(query, userId, symbol).Scan(
		&pnl.ID,
		&pnl.UserID,
		&pnl.AssetSymbol,
		&pnl.AssetType,
		&pnl.Quantity,
		&pnl.AverageCostUSD,
		&pnl.AverageCostBRL,
		&pnl.TotalCostUSD,
		&pnl.TotalCostBRL,
		&pnl.SellingPriceUSD,
		&pnl.SellingPriceBRL,
		&pnl.TotalValueSoldUSD,
		&pnl.TotalValueSoldBRL,
		&pnl.RealizedProfitUSD,
		&pnl.RealizedProfitBRL,
		&pnl.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &pnl, nil
}

func (r *realizedPNLRepository) GetPNLs(userId string) ([]domain.RealizedPNL, error) {
	query := `
		SELECT 
			id,
			user_id,
			asset_symbol,
			asset_type,
			quantity,
			average_cost_usd,
			average_cost_brl,
			total_cost_usd,
			total_cost_brl,
			selling_price_usd,
			selling_price_brl,
			total_value_sold_usd,
			total_value_sold_brl,
			realized_profit_usd,
			realized_profit_brl,
			created_at
		FROM realized_pnl
		WHERE user_id = $1
	`

	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pnls []domain.RealizedPNL
	for rows.Next() {
		var pnl domain.RealizedPNL
		err := rows.Scan(
			&pnl.ID,
			&pnl.UserID,
			&pnl.AssetSymbol,
			&pnl.AssetType,
			&pnl.Quantity,
			&pnl.AverageCostUSD,
			&pnl.AverageCostBRL,
			&pnl.TotalCostUSD,
			&pnl.TotalCostBRL,
			&pnl.SellingPriceUSD,
			&pnl.SellingPriceBRL,
			&pnl.TotalValueSoldUSD,
			&pnl.TotalValueSoldBRL,
			&pnl.RealizedProfitUSD,
			&pnl.RealizedProfitBRL,
			&pnl.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		pnls = append(pnls, pnl)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return pnls, nil
}
