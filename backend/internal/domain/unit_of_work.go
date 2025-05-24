package domain

// UnitOfWork interface para gerenciar transações distribuídas
type UnitOfWork interface {
	Begin() error
	Commit() error
	Rollback() error
	GetTransactionRepository() TransactionRepository
	GetPositionRepository() PositionRepository
}
