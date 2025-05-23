package domain

import "time"

type CurrencyService interface {
	GetUSDToBRL(date time.Time) (float64, error)
}