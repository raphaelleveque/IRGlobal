package realizedpnl

import (
	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
)

type realizedPNLService struct {
	repo domain.RealizedPNLRepository
}

func NewRealizedPNLService(repo domain.RealizedPNLRepository) domain.RealizedPNLService {
	return &realizedPNLService{repo: repo}
}

func (s *realizedPNLService) CalculatePNL(transaction *domain.Transaction, position *domain.Position, dbTx domain.DBTx) (*domain.RealizedPNL, error) {
	pnl, err := s.repo.GetPNLByAssetSymbol(transaction.UserID, transaction.AssetSymbol)
	if err != nil {
		return nil, err
	}
	if pnl == nil {
		pnl = &domain.RealizedPNL{
			UserID:      transaction.UserID,
			AssetSymbol: transaction.AssetSymbol,
			AssetType:   transaction.AssetType,
		}
	}

	// Calcular usando a quantidade anterior
	pnl.AverageCostUSD, pnl.AverageCostBRL = s.setAverageCost(position)
	pnl.SellingPriceUSD, pnl.SellingPriceBRL = s.setSellingPrice(transaction, pnl)

	// Atualizar a quantidade após os cálculos de média
	pnl.Quantity = s.setQuantity(transaction, pnl)

	// Calcular valores totais com a nova quantidade
	pnl.TotalCostUSD, pnl.TotalCostBRL = s.setTotalCost(pnl)
	pnl.TotalValueSoldUSD, pnl.TotalValueSoldBRL = s.setTotalValueSold(pnl)
	pnl.RealizedProfitUSD, pnl.RealizedProfitBRL = s.setRealizedProfit(pnl)

	// Salvar o PNL no banco
	pnl, err = s.repo.UpdatePNL(pnl, dbTx)
	if err != nil {
		return nil, err
	}

	return pnl, nil
}

func (s *realizedPNLService) RecalculatePNL(userId, symbol string, dbTx domain.DBTx) (*domain.RealizedPNL, error) {
	return nil, nil
}

func (s *realizedPNLService) setQuantity(transaction *domain.Transaction, pnl *domain.RealizedPNL) float64 {
	return pnl.Quantity + transaction.Quantity
}

func (s *realizedPNLService) setAverageCost(position *domain.Position) (AverageCostUSD float64, AverageCostBRL float64) {
	return position.AverageCostUSD, position.AverageCostBRL
}

func (s *realizedPNLService) setTotalCost(pnl *domain.RealizedPNL) (TotalCostUSD float64, TotalCostBRL float64) {
	return pnl.AverageCostUSD * pnl.Quantity, pnl.AverageCostBRL * pnl.Quantity
}

func (s *realizedPNLService) setSellingPrice(transaction *domain.Transaction, pnl *domain.RealizedPNL) (SellingPriceUSD float64, SellingPriceBRL float64) {
	if pnl.Quantity == 0 {
		return transaction.PriceInUSD, transaction.PriceInBRL
	}

	sellingPriceUSD := ((pnl.SellingPriceUSD * pnl.Quantity) + (transaction.PriceInUSD * transaction.Quantity)) / (pnl.Quantity + transaction.Quantity)
	sellingPriceBRL := ((pnl.SellingPriceBRL * pnl.Quantity) + (transaction.PriceInBRL * transaction.Quantity)) / (pnl.Quantity + transaction.Quantity)

	return sellingPriceUSD, sellingPriceBRL
}

func (s *realizedPNLService) setTotalValueSold(pnl *domain.RealizedPNL) (TotalValueSoldUSD float64, TotalValueSoldBRL float64) {
	totalValueSoldUSD := pnl.SellingPriceUSD * pnl.Quantity
	totalValueSoldBRL := pnl.SellingPriceBRL * pnl.Quantity

	return totalValueSoldUSD, totalValueSoldBRL
}

func (s *realizedPNLService) setRealizedProfit(pnl *domain.RealizedPNL) (RealizedProfitUSD float64, RealizedProfitBRL float64) {
	realizedProfitUSD := pnl.TotalValueSoldUSD - pnl.TotalCostUSD
	realizedProfitBRL := pnl.TotalValueSoldBRL - pnl.TotalCostBRL

	return realizedProfitUSD, realizedProfitBRL
}
