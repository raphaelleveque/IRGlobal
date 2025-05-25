package realizedpnl

import (
	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
)

type realizedPNLService struct {
	repo               domain.RealizedPNLRepository
	transactionService domain.TransactionService
}

func NewRealizedPNLService(repo domain.RealizedPNLRepository, transactionService domain.TransactionService) domain.RealizedPNLService {
	return &realizedPNLService{
		repo:               repo,
		transactionService: transactionService,
	}
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
	allTransactions, err := s.transactionService.FindAllBySymbol(userId, symbol)
	if err != nil {
		return nil, err
	}

	sellTransactions := s.filterSellTransactions(allTransactions)

	if len(sellTransactions) == 0 {
		_, err := s.repo.DeletePNL(userId, symbol, dbTx)
		return nil, err
	}

	pnl := &domain.RealizedPNL{
		UserID:      userId,
		AssetSymbol: symbol,
		AssetType:   sellTransactions[0].AssetType,
	}

	for _, sellTransaction := range sellTransactions {
		pnl.AverageCostUSD, pnl.AverageCostBRL = s.calculateAverageCostAtSell(allTransactions, sellTransaction)
		pnl.SellingPriceUSD, pnl.SellingPriceBRL = s.setSellingPrice(&sellTransaction, pnl)
		pnl.Quantity = s.setQuantity(&sellTransaction, pnl)
		pnl = s.recalculateTotals(pnl)
	}

	return s.repo.UpdatePNL(pnl, dbTx)
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

// calculateAverageCostAtSell calcula o custo médio da posição no momento de uma venda específica
// Processa apenas as compras (BUY) que ocorreram antes desta venda
func (s *realizedPNLService) calculateAverageCostAtSell(allTransactions []domain.Transaction, sellTransaction domain.Transaction) (float64, float64) {
	var totalQuantity float64
	var totalCostUSD float64
	var totalCostBRL float64

	for _, transaction := range allTransactions {
		if transaction.Type == domain.Buy && !transaction.OperationDate.After(sellTransaction.OperationDate) {
			totalQuantity += transaction.Quantity
			totalCostUSD += transaction.TotalCostUSD
			totalCostBRL += transaction.TotalCostBRL
		}
	}

	if totalQuantity == 0 {
		return 0, 0
	}

	avgCostUSD := totalCostUSD / totalQuantity
	avgCostBRL := totalCostBRL / totalQuantity

	return avgCostUSD, avgCostBRL
}

// filterSellTransactions filtra apenas as transações de venda
func (s *realizedPNLService) filterSellTransactions(transactions []domain.Transaction) []domain.Transaction {
	var sellTransactions []domain.Transaction
	for _, transaction := range transactions {
		if transaction.Type == domain.Sell {
			sellTransactions = append(sellTransactions, transaction)
		}
	}
	return sellTransactions
}

// recalculateTotals recalcula todos os valores totais do PNL
func (s *realizedPNLService) recalculateTotals(pnl *domain.RealizedPNL) *domain.RealizedPNL {
	pnl.TotalCostUSD, pnl.TotalCostBRL = s.setTotalCost(pnl)
	pnl.TotalValueSoldUSD, pnl.TotalValueSoldBRL = s.setTotalValueSold(pnl)
	pnl.RealizedProfitUSD, pnl.RealizedProfitBRL = s.setRealizedProfit(pnl)
	return pnl
}
