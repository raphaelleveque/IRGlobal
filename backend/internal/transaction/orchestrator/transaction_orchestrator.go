package orchestrator

import "github.com/raphaelleveque/IRGlobal/backend/internal/domain"

type TransactionOrchestrator struct {
	transactionService domain.TransactionService
	positionService    domain.PositionService
	pnlService domain.RealizedPNLService
	uow                domain.UnitOfWork
}

func NewTransactionOrchestrator(transactionService domain.TransactionService, positionService domain.PositionService, pnlService domain.RealizedPNLService, uow domain.UnitOfWork) *TransactionOrchestrator {
	return &TransactionOrchestrator{
		transactionService: transactionService,
		positionService:    positionService,
		pnlService: pnlService,
		uow:                uow,
	}
}

func (o *TransactionOrchestrator) AddTransactionWithPosition(transaction *domain.Transaction) (*domain.Transaction, *domain.Position, *domain.RealizedPNL, error) {
	dbTx, err := o.uow.Begin()
	if err != nil {
		return nil, nil, nil, err
	}
	defer func() {
		if err != nil {
			dbTx.Rollback()
		}
	}()

	// Add transaction
	transaction, err = o.transactionService.AddTransaction(transaction, dbTx)
	if err != nil {
		return nil, nil, nil, err
	}

	// Calculate position
	position, err := o.positionService.CalculatePosition(transaction, dbTx)
	if err != nil {
		return nil, nil, nil, err
	}

	var realizedPnl *domain.RealizedPNL
	if transaction.Type == domain.Sell {
		realizedPnl, err = o.pnlService.CalculatePNL(transaction, position, dbTx)
		if err != nil {
			return nil, nil, nil, err
		}
	}

	// Commit transaction
	if err = dbTx.Commit(); err != nil {
		return nil, nil, nil, err
	}

	return transaction, position, realizedPnl, nil
}

func (o *TransactionOrchestrator) DeleteTransactionWithPosition(transactionID string, userID string) (*domain.Transaction, *domain.Position, error) {
	dbTx, err := o.uow.Begin()
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		if err != nil {
			dbTx.Rollback()
		}
	}()

	// Delete transaction
	transaction, err := o.transactionService.DeleteTransaction(transactionID, dbTx)
	if err != nil {
		return nil, nil, err
	}

	// Recalculate position
	position, err := o.positionService.RecalculatePosition(userID, transaction.AssetSymbol, dbTx)
	if err != nil {
		return nil, nil, err
	}

	// Commit transaction
	if err = dbTx.Commit(); err != nil {
		return nil, nil, err
	}

	return transaction, position, nil
}
