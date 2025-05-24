package transaction

import (
	"database/sql"

	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
)

type transactionCoordinator struct {
	db                 *sql.DB
	transactionService domain.TransactionService
	positionService    domain.PositionService
}

func NewTransactionCoordinator(db *sql.DB, transactionService domain.TransactionService, positionService domain.PositionService, currencyService domain.CurrencyService) domain.TransactionCoordinatorService {
	return &transactionCoordinator{
		db:                 db,
		transactionService: transactionService,
		positionService:    positionService,
	}
}

func (tc *transactionCoordinator) AddTransactionWithPosition(transaction *domain.Transaction) (*domain.Transaction, *domain.Position, error) {
	var createdTransaction *domain.Transaction
	var updatedPosition *domain.Position

	err := tc.withTransaction(func(tx *sql.Tx) error {
		// Criar repositories transacionais
		transactionRepo := NewTransactionRepositoryWithTx(tx)
		positionRepo := newPositionRepositoryWithTx(tx)

		// Preparar transação (cálculos de preço e câmbio) usando TransactionService
		err := tc.transactionService.PrepareTransaction(transaction)
		if err != nil {
			return err
		}

		// Validar transação usando PositionService
		err = tc.positionService.ValidateTransaction(transaction, positionRepo)
		if err != nil {
			return err
		}

		// Criar transação
		createdTransaction, err = transactionRepo.Create(transaction)
		if err != nil {
			return err
		}

		// Calcular e salvar posição usando PositionService
		updatedPosition, err = tc.positionService.CalculatePositionWithRepo(createdTransaction, positionRepo)
		if err != nil {
			return err
		}

		return nil
	})

	return createdTransaction, updatedPosition, err
}

func (tc *transactionCoordinator) DeleteTransactionWithPosition(id string) (*domain.Transaction, *domain.Position, error) {
	var deletedTransaction *domain.Transaction
	var updatedPosition *domain.Position

	err := tc.withTransaction(func(tx *sql.Tx) error {
		// Criar repositories transacionais
		transactionRepo := NewTransactionRepositoryWithTx(tx)
		positionRepo := newPositionRepositoryWithTx(tx)

		// Deletar transação
		var err error
		deletedTransaction, err = transactionRepo.Delete(id)
		if err != nil {
			return err
		}

		// Recalcular posição usando PositionService
		updatedPosition, err = tc.positionService.RecalculatePositionWithRepo(deletedTransaction.UserID, deletedTransaction.AssetSymbol, positionRepo, transactionRepo)
		if err != nil {
			return err
		}

		return nil
	})

	return deletedTransaction, updatedPosition, err
}

// Função local para criar position repository (evita import cycle)
func newPositionRepositoryWithTx(tx *sql.Tx) domain.PositionRepository {
	return &positionRepository{db: tx}
}

// Implementação local do position repository para evitar import cycle
type positionRepository struct {
	db dbExecutor
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
		ON CONFLICT (user_id, asset_symbol) 
		DO UPDATE SET 
			quantity = EXCLUDED.quantity,
			average_cost_usd = EXCLUDED.average_cost_usd,
			average_cost_brl = EXCLUDED.average_cost_brl,
			total_cost_usd = EXCLUDED.total_cost_usd,
			total_cost_brl = EXCLUDED.total_cost_brl
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

// Método auxiliar para gerenciar transações
func (tc *transactionCoordinator) withTransaction(fn func(tx *sql.Tx) error) error {
	tx, err := tc.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
