package position

import (
	"errors"

	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
)

type positionService struct {
	repo               domain.PositionRepository
	transactionService domain.TransactionService
}

func NewPositionService(repo domain.PositionRepository, transactionService domain.TransactionService) domain.PositionService {
	return &positionService{repo: repo, transactionService: transactionService}
}

func (s *positionService) CalculatePosition(transaction *domain.Transaction, dbTx domain.DBTx) (*domain.Position, error) {
	position, err := s.repo.GetPositionByAssetSymbol(transaction.UserID, transaction.AssetSymbol)
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
	if positionQuantity == 0 {
		_, err = s.repo.DeletePosition(position.UserID, position.AssetSymbol, dbTx)
		return nil, err
	}

	position.Quantity = positionQuantity
	position.TotalCostUSD, position.TotalCostBRL = s.calculateTotalCost(position)

	position, err = s.repo.UpdatePosition(position, dbTx)
	if err != nil {
		return nil, err
	}
	return position, nil
}

func (s *positionService) RecalculatePosition(userId, symbol, transactionId string, dbTx domain.DBTx) (*domain.Position, error) {
	transactions, err := s.transactionService.FindAllBySymbolExcludingOne(userId, symbol, transactionId)
	if err != nil {
		return nil, err
	}

	// If there are no transactions, delete the position
	if len(transactions) == 0 {
		_, err = s.repo.DeletePosition(userId, symbol, dbTx)
		return nil, err
	}

	// Recalculate position from scratch
	position := &domain.Position{
		UserID:      userId,
		AssetSymbol: symbol,
		AssetType:   transactions[0].AssetType, // Use the type from the first transaction
	}

	// Process all transactions in order
	for _, transaction := range transactions {
		position.AverageCostUSD, position.AverageCostBRL = s.calculateAverageCost(&transaction, position)
		positionQuantity, err := s.calculateNewQuantity(&transaction, position)
		if err != nil {
			return nil, err
		}
		position.Quantity = positionQuantity
		position.TotalCostUSD, position.TotalCostBRL = s.calculateTotalCost(position)
	}

	if position.Quantity == 0 {
		_, err = s.repo.DeletePosition(position.UserID, position.AssetSymbol, dbTx)
		return nil, err
	}

	position, err = s.repo.UpdatePosition(position, dbTx)
	if err != nil {
		return nil, err
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
		return positionQuantity, errors.New("you cannot sell more quantities than what you have")
	}

	return positionQuantity, nil
}

func (s *positionService) calculateAverageCost(transaction *domain.Transaction, position *domain.Position) (averageCostUSD float64, averageCostBRL float64) {
	if transaction.Type == domain.Sell {
		return position.AverageCostUSD, position.AverageCostBRL
	}

	newAvgCostUsd := ((position.Quantity * position.AverageCostUSD) + (transaction.Quantity * transaction.PriceInUSD)) / (position.Quantity + transaction.Quantity)
	newAvgCostBrl := ((position.Quantity * position.AverageCostBRL) + (transaction.Quantity * transaction.PriceInBRL)) / (position.Quantity + transaction.Quantity)

	return newAvgCostUsd, newAvgCostBrl

}

func (s *positionService) calculateTotalCost(position *domain.Position) (totalCostUSD float64, totalCostBRL float64) {
	return (position.AverageCostUSD * position.Quantity), (position.AverageCostBRL * position.Quantity)
}

func (s *positionService) GetPositionByAssetSymbol(userId, symbol string) (*domain.Position, error) {
	return s.repo.GetPositionByAssetSymbol(userId, symbol)
}

func (s *positionService) GetPositions(userId string) ([]domain.Position, error) {
	return s.repo.GetPositions(userId)
}
