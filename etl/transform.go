package etl

import (
	"fmt"
	"strconv"

	"bitdeal.nl/models"
)

func Transform(message models.GetPricesData, t interface{}) models.ExchangeRate {
	switch exchangeData := t.(type) {
	case *models.Bitvavo:
		rate, _ := strconv.ParseFloat(exchangeData.Asks[0][0], 32)
		amount := message.Amount * rate
		return models.ExchangeRate{
			Exchange: "Bitvavo",
			Rate:     rate,
			Amount:   amount,
		}
	case *models.Litebit:
		fmt.Println("Litebit")
		return models.ExchangeRate{}
	}

	return models.ExchangeRate{}
}
