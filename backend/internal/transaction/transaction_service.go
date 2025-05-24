package transaction

import (
	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
)

type transactionService struct {
	repo            domain.TransactionRepository
	currencyService domain.CurrencyService
	positionService domain.PositionService
}

func NewTransactionService(repo domain.TransactionRepository, currencyService domain.CurrencyService, positionService domain.PositionService) domain.TransactionService {
	return &transactionService{
		repo:            repo,
		currencyService: currencyService,
		positionService: positionService,
	}
}

func (s *transactionService) AddTransaction(transaction *domain.Transaction) (*domain.Transaction, error) {
	usdbrlRate, err := s.currencyService.GetUSDToBRL(transaction.OperationDate)
	if err != nil {
		return transaction, err
	}

	s.setBRLPrice(transaction, usdbrlRate)
	s.setTotalCost(transaction)

	transaction, err = s.repo.Create(transaction)
	return transaction, err
}

func (s *transactionService) DeleteTransaction(id string) (*domain.Transaction, error) {
	transaction, err := s.repo.Delete(id)
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

// Métodos transacionais
func (s *transactionService) AddTransactionWithPosition(transaction *domain.Transaction, uow domain.UnitOfWork) (*domain.Transaction, *domain.Position, error) {
	// Começar a transação
	if err := uow.Begin(); err != nil {
		return nil, nil, err
	}

	// Calcular taxa de câmbio e preços
	usdbrlRate, err := s.currencyService.GetUSDToBRL(transaction.OperationDate)
	if err != nil {
		uow.Rollback()
		return nil, nil, err
	}

	s.setBRLPrice(transaction, usdbrlRate)
	s.setTotalCost(transaction)

	// Criar transação usando o repository transacional
	transactionRepo := uow.GetTransactionRepository()
	createdTransaction, err := transactionRepo.Create(transaction)
	if err != nil {
		uow.Rollback()
		return nil, nil, err
	}

	// Calcular e salvar posição usando o repository transacional
	positionRepo := uow.GetPositionRepository()
	position, err := s.calculateAndSavePosition(createdTransaction, positionRepo)
	if err != nil {
		uow.Rollback()
		return nil, nil, err
	}

	// Commit da transação
	if err := uow.Commit(); err != nil {
		return nil, nil, err
	}

	return createdTransaction, position, nil
}

func (s *transactionService) DeleteTransactionWithPosition(id string, uow domain.UnitOfWork) (*domain.Transaction, *domain.Position, error) {
	// Começar a transação
	if err := uow.Begin(); err != nil {
		return nil, nil, err
	}

	// Buscar e deletar transação usando o repository transacional
	transactionRepo := uow.GetTransactionRepository()
	deletedTransaction, err := transactionRepo.Delete(id)
	if err != nil {
		uow.Rollback()
		return nil, nil, err
	}

	// Recalcular posição
	positionRepo := uow.GetPositionRepository()
	position, err := s.recalculatePosition(deletedTransaction.UserID, deletedTransaction.AssetSymbol, transactionRepo, positionRepo)
	if err != nil {
		uow.Rollback()
		return nil, nil, err
	}

	// Commit da transação
	if err := uow.Commit(); err != nil {
		return nil, nil, err
	}

	return deletedTransaction, position, nil
}

// Métodos auxiliares
func (s *transactionService) setBRLPrice(transaction *domain.Transaction, usdbrlRate float64) {
	transaction.USDBRLRate = usdbrlRate
	transaction.PriceInBRL = transaction.PriceInUSD * transaction.USDBRLRate
}

func (s *transactionService) setTotalCost(transaction *domain.Transaction) {
	transaction.TotalCostUSD = transaction.PriceInUSD * transaction.Quantity
	transaction.TotalCostBRL = transaction.PriceInBRL * transaction.Quantity
}

func (s *transactionService) calculateAndSavePosition(transaction *domain.Transaction, positionRepo domain.PositionRepository) (*domain.Position, error) {
	position, err := positionRepo.GetPositionByAssetSymbol(transaction.UserID, transaction.AssetSymbol)
	if err != nil {
		return nil, err
	}

	if position == nil {
		position = &domain.Position{
			UserID:      transaction.UserID,
			AssetSymbol: transaction.AssetSymbol,
			AssetType:   transaction.AssetType,
		}
	}

	position.AverageCostUSD, position.AverageCostBRL = s.calculateAverageCost(transaction, position)
	positionQuantity, err := s.calculateNewQuantity(transaction, position)
	if err != nil {
		return nil, err
	}
	position.Quantity = positionQuantity
	position.TotalCostUSD, position.TotalCostBRL = s.calculateTotalCost(position)

	return positionRepo.UpdatePosition(position)
}

func (s *transactionService) recalculatePosition(userId, symbol string, transactionRepo domain.TransactionRepository, positionRepo domain.PositionRepository) (*domain.Position, error) {
	transactions, err := transactionRepo.FindAllBySymbol(userId, symbol)
	if err != nil {
		return nil, err
	}

	// Deletar posição atual
	_, err = positionRepo.DeletePosition(userId, symbol)
	if err != nil {
		return nil, err
	}

	var position *domain.Position
	for _, transaction := range transactions {
		position, err = s.calculateAndSavePosition(&transaction, positionRepo)
		if err != nil {
			return nil, err
		}
	}
	return position, nil
}

func (s *transactionService) calculateNewQuantity(transaction *domain.Transaction, position *domain.Position) (float64, error) {
	var positionQuantity float64
	switch transaction.Type {
	case domain.Buy:
		positionQuantity = position.Quantity + transaction.Quantity
	case domain.Sell:
		positionQuantity = position.Quantity - transaction.Quantity
	default:
		positionQuantity = 0.0
	}

	if positionQuantity < 0 {
		return positionQuantity, domain.ErrInsufficientQuantity
	}

	return positionQuantity, nil
}

func (s *transactionService) calculateAverageCost(transaction *domain.Transaction, position *domain.Position) (averageCostUSD float64, averageCostBRL float64) {
	if transaction.Type == domain.Sell {
		return position.AverageCostUSD, position.AverageCostBRL
	}

	newAvgCostUsd := ((position.Quantity * position.AverageCostUSD) + (transaction.Quantity * transaction.PriceInUSD)) / (position.Quantity + transaction.Quantity)
	newAvgCostBrl := ((position.Quantity * position.AverageCostBRL) + (transaction.Quantity * transaction.PriceInBRL)) / (position.Quantity + transaction.Quantity)

	return newAvgCostUsd, newAvgCostBrl
}

func (s *transactionService) calculateTotalCost(position *domain.Position) (totalCostUSD float64, totalCostBRL float64) {
	return (position.AverageCostUSD * position.Quantity), (position.AverageCostBRL * position.Quantity)
}
