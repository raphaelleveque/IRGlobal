package currency

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

func FetchUSDToBRL(date time.Time) (float64, error) {
	url := fmt.Sprintf(
		"https://olinda.bcb.gov.br/olinda/servico/PTAX/versao/v1/odata/CotacaoDolarDia(dataCotacao=@dataCotacao)?@dataCotacao='%s'&$top=1&$format=json",
		date.Format("01-02-2006"),
	)

	log.Println("Requisição HTTP para URL:", url)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result Response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	responseJSON, _ := json.Marshal(result)
	log.Println("Resposta JSON:", string(responseJSON))

	if len(result.Value) == 0 {
		return 0, errors.New("USD BRL Rate not found")
	}

	return result.Value[0].BuyRate, nil
}
