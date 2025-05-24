package position

import (
	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
)

type positionService struct {
	repo               domain.PositionRepository
	transactionService domain.TransactionService
}

func NewPositionService(repo domain.PositionRepository, transactionService domain.TransactionService) domain.PositionService {
	return &positionService{
		repo:               repo,
		transactionService: transactionService,
	}
}

func (s *positionService) CalculatePosition(transaction *domain.Transaction) (*domain.Position, error) {
	return s.calculatePositionWithRepo(transaction, s.repo)
}

func (s *positionService) RecalculatePosition(userId, symbol string) (*domain.Position, error) {
	return s.recalculatePositionWithRepo(userId, symbol, s.repo, s.transactionService)
}

// Métodos que permitem usar repositories específicos (para transações)
func (s *positionService) CalculatePositionWithRepo(transaction *domain.Transaction, positionRepo domain.PositionRepository) (*domain.Position, error) {
	return s.calculatePositionWithRepo(transaction, positionRepo)
}

func (s *positionService) RecalculatePositionWithRepo(userId, symbol string, positionRepo domain.PositionRepository, transactionRepo domain.TransactionRepository) (*domain.Position, error) {
	return s.recalculatePositionWithRepo(userId, symbol, positionRepo, transactionRepo)
}

func (s *positionService) ValidateTransaction(transaction *domain.Transaction, positionRepo domain.PositionRepository) error {
	position, err := positionRepo.GetPositionByAssetSymbol(transaction.UserID, transaction.AssetSymbol)
	if err != nil {
		return err
	}

	if position == nil {
		position = &domain.Position{
			UserID:      transaction.UserID,
			AssetSymbol: transaction.AssetSymbol,
			AssetType:   transaction.AssetType,
		}
	}

	// Validar se é possível fazer a operação
	_, err = s.calculateNewQuantity(transaction, position)
	return err
}

// Métodos internos que implementam a lógica de negócio
func (s *positionService) calculatePositionWithRepo(transaction *domain.Transaction, positionRepo domain.PositionRepository) (*domain.Position, error) {
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

	position, err = positionRepo.UpdatePosition(position)
	if err != nil {
		return nil, err
	}
	return position, nil
}

func (s *positionService) recalculatePositionWithRepo(userId, symbol string, positionRepo domain.PositionRepository, transactionService interface{}) (*domain.Position, error) {
	// Usar o transactionService passado ou o padrão
	var transactions []domain.Transaction
	var err error

	if txRepo, ok := transactionService.(domain.TransactionRepository); ok {
		// Se for um repository transacional, usar ele
		transactions, err = txRepo.FindAllBySymbol(userId, symbol)
	} else {
		// Senão, usar o service padrão
		transactions, err = s.transactionService.FindAllBySymbol(userId, symbol)
	}

	if err != nil {
		return nil, err
	}

	_, err = positionRepo.DeletePosition(userId, symbol)
	if err != nil {
		return nil, err
	}

	var position *domain.Position
	for _, transaction := range transactions {
		position, err = s.calculatePositionWithRepo(&transaction, positionRepo)
		if err != nil {
			return nil, err
		}
	}
	return position, nil
}

func (s *positionService) calculateNewQuantity(transaction *domain.Transaction, position *domain.Position) (float64, error) {
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

func (s *positionService) calculateAverageCost(transaction *domain.Transaction, position *domain.Position) (averageCostUSD float64, averageCostBRL float64) {
	if transaction.Type == domain.Sell {
		return position.AverageCostUSD, position.AverageCostBRL
	}

	if position.Quantity == 0 {
		return transaction.PriceInUSD, transaction.PriceInBRL
	}

	newAvgCostUsd := ((position.Quantity * position.AverageCostUSD) + (transaction.Quantity * transaction.PriceInUSD)) / (position.Quantity + transaction.Quantity)
	newAvgCostBrl := ((position.Quantity * position.AverageCostBRL) + (transaction.Quantity * transaction.PriceInBRL)) / (position.Quantity + transaction.Quantity)

	return newAvgCostUsd, newAvgCostBrl
}

func (s *positionService) calculateTotalCost(position *domain.Position) (totalCostUSD float64, totalCostBRL float64) {
	return (position.AverageCostUSD * position.Quantity), (position.AverageCostBRL * position.Quantity)
}
