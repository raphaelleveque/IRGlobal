package infrastructure

import (
	"database/sql"

	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
)

type SqlUnitOfWork struct {
	DB *sql.DB
}

type SqlDBTx struct {
	Tx *sql.Tx
}

func NewSqlUnitOfWork(db *sql.DB) *SqlUnitOfWork {
	return &SqlUnitOfWork{DB: db}
}

func (s *SqlUnitOfWork) Begin() (domain.DBTx, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, err
	}

	return &SqlDBTx{Tx: tx}, nil
}

func (s *SqlDBTx) Commit() error {
	return s.Tx.Commit()
}

func (s *SqlDBTx) Rollback() error {
	return s.Tx.Rollback()
}
