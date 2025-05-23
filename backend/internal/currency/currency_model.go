package currency

import "time"

type Response struct {
	ODataContext string  `json:"@odata.context"`
	Value        []Quote `json:"value"`
}

type Quote struct {
	BuyRate  float64 `json:"cotacaoCompra"`
	SellRate float64 `json:"cotacaoVenda"`
	RateDate string  `json:"dataHoraCotacao"`
}

// ParseRateDate converte a string de data para time.Time quando necessário
func (q *Quote) ParseRateDate() (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05.999", q.RateDate)
}
