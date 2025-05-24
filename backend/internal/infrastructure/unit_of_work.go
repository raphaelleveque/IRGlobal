package infrastructure

import (
	"database/sql"

	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
	"github.com/raphaelleveque/IRGlobal/backend/internal/position"
	"github.com/raphaelleveque/IRGlobal/backend/internal/transaction"
)

type unitOfWork struct {
	db                    *sql.DB
	tx                    *sql.Tx
	transactionRepository domain.TransactionRepository
	positionRepository    domain.PositionRepository
}

func NewUnitOfWork(db *sql.DB) domain.UnitOfWork {
	return &unitOfWork{
		db: db,
	}
}

func (uow *unitOfWork) Begin() error {
	tx, err := uow.db.Begin()
	if err != nil {
		return err
	}
	uow.tx = tx

	// Criar repositories que usam a transação
	uow.transactionRepository = transaction.NewTransactionRepositoryWithTx(tx)
	uow.positionRepository = position.NewPositionRepositoryWithTx(tx)

	return nil
}

func (uow *unitOfWork) Commit() error {
	if uow.tx == nil {
		return nil
	}
	return uow.tx.Commit()
}

func (uow *unitOfWork) Rollback() error {
	if uow.tx == nil {
		return nil
	}
	return uow.tx.Rollback()
}

func (uow *unitOfWork) GetTransactionRepository() domain.TransactionRepository {
	return uow.transactionRepository
}

func (uow *unitOfWork) GetPositionRepository() domain.PositionRepository {
	return uow.positionRepository
}
