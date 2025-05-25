package transaction

import (
	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
)

type transactionService struct {
	repo            domain.TransactionRepository
	currencyService domain.CurrencyService
}

func NewTransactionService(repo domain.TransactionRepository, currencyService domain.CurrencyService) domain.TransactionService {
	return &transactionService{repo: repo, currencyService: currencyService}
}

func (s *transactionService) AddTransaction(transaction *domain.Transaction, dbTx domain.DBTx) (*domain.Transaction, error) {
	usdbrlRate, err := s.currencyService.GetUSDToBRL(transaction.OperationDate)
	if err != nil {
		return transaction, err
	}

	s.setBRLPrice(transaction, usdbrlRate)
	s.setTotalCost(transaction)

	transaction, err = s.repo.Create(transaction, dbTx)
	return transaction, err
}

func (s *transactionService) DeleteTransaction(id string, dbTx domain.DBTx) (*domain.Transaction, error) {
	transaction, err := s.repo.Delete(id, dbTx)
	return transaction, err
}

func (s *transactionService) FindByID(id string) (*domain.Transaction, error) {
	transaction, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (s *transactionService) FindAllBySymbol(userId, symbol string) ([]domain.Transaction, error) {
	return s.repo.FindAllBySymbol(userId, symbol)
}

func (s *transactionService) FindAllBySymbolExcludingOne(userId, symbol, transactionId string) ([]domain.Transaction, error) {
	return s.repo.FindAllBySymbolExcludingOne(userId, symbol, transactionId)
}

func (s *transactionService) setBRLPrice(transaction *domain.Transaction, usdbrlRate float64) {
	transaction.USDBRLRate = usdbrlRate
	transaction.PriceInBRL = transaction.PriceInUSD * transaction.USDBRLRate
}

func (s *transactionService) setTotalCost(transaction *domain.Transaction) {
	transaction.TotalCostUSD = transaction.PriceInUSD * transaction.Quantity
	transaction.TotalCostBRL = transaction.PriceInBRL * transaction.Quantity
}
