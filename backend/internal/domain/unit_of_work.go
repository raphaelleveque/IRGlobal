package domain

type DBTx interface {
	Commit() error
	Rollback() error
}

type UnitOfWork interface {
	Begin() (DBTx, error)
}