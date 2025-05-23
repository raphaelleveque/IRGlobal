package currency

import (
	"time"

	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
)

type currencyService struct {
}

func NewCurrencyService() domain.CurrencyService {
	return &currencyService{}
}

func (c *currencyService) GetUSDToBRL(date time.Time) (float64, error) {
	return FetchUSDToBRL(date)
}