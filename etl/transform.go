package etl

import (
	"math"
	"strconv"

	"bitdeal.nl/models"
)

/*
Transform ...
*/
func Transform(request models.GetPricesData, t interface{}) models.ExchangeRate {
	var rate float64
	var amount float64

	switch exchangeData := t.(type) {
	case *models.Bitvavo:
		if request.Type == "buy" {
			rate, _ = strconv.ParseFloat(exchangeData.Asks[0][0], 32)
		} else {
			rate, _ = strconv.ParseFloat(exchangeData.Bids[0][0], 32)
		}

		reqAmount, _ := strconv.ParseFloat(request.Amount, 64)
		if request.Currency == "btc" {
			amount = math.Round(reqAmount*rate*100) / 100
		} else {
			amount = math.Round(reqAmount/rate*100000000) / 100000000
		}

		return models.ExchangeRate{
			Exchange: "Bitvavo",
			Rate:     math.Round(rate*100) / 100,
			Amount:   amount,
			Link:     "https://bitvavo.com/?a=C23D4F6D97",
			Reviews:  1895,
			Stars:    5,
		}
	case *models.Litebit:
		if request.Type == "buy" {
			rate, _ = strconv.ParseFloat(exchangeData.Result.Buy, 32)
		} else {
			rate, _ = strconv.ParseFloat(exchangeData.Result.Sell, 32)
		}

		reqAmount, _ := strconv.ParseFloat(request.Amount, 64)
		if request.Currency == "btc" {
			amount = math.Round(reqAmount*rate*100) / 100
		} else {
			amount = math.Round(reqAmount/rate*100000000) / 100000000
		}

		return models.ExchangeRate{
			Exchange: "Litebit",
			Rate:     math.Round(rate*100) / 100,
			Amount:   amount,
			Link:     "https://www.litebit.eu/?referrer=40203",
			Reviews:  1695,
			Stars:    5,
		}
	case *models.BitonicBuy:
		if request.Currency == "btc" {
			amount = math.Round(exchangeData.Eur*100) / 100
		} else {
			amount = exchangeData.Btc
		}

		return models.ExchangeRate{
			Exchange: "Bitonic",
			Rate:     math.Round(exchangeData.Price*100) / 100,
			Amount:   amount,
			Link:     "https://bitonic.nl/?partner=777",
			Reviews:  41,
			Stars:    4,
		}
	case *models.BitonicSell:
		if request.Currency == "btc" {
			amount = math.Round(exchangeData.Eur*100) / 100
		} else {
			amount = exchangeData.Btc
		}

		return models.ExchangeRate{
			Exchange: "Bitonic",
			Rate:     math.Round(exchangeData.Price*100) / 100,
			Amount:   amount,
			Link:     "https://bitonic.nl/?partner=777",
			Reviews:  41,
			Stars:    4,
		}
	case *models.Btcdirect:
		if request.Type == "buy" {
			rate, _ = strconv.ParseFloat(exchangeData.ExchangeRates.Sell.Btc, 32)
		} else {
			rate, _ = strconv.ParseFloat(exchangeData.ExchangeRates.Buy.Btc, 32)
		}

		reqAmount, _ := strconv.ParseFloat(request.Amount, 64)
		if request.Currency == "btc" {
			amount = math.Round(reqAmount*rate*100) / 100
		} else {
			amount = math.Round(reqAmount/rate*100000000) / 100000000
		}

		return models.ExchangeRate{
			Exchange: "Btcdirect",
			Rate:     math.Round(rate*100) / 100,
			Amount:   amount,
			Link:     "https://btcdirect.eu/nl-nl?partnerId=319",
			Reviews:  5485,
			Stars:    5,
		}
	case *models.BitrushBuy:
		reqAmount, _ := strconv.ParseFloat(request.Amount, 64)
		if request.Currency == "btc" {
			rate = exchangeData.Amount / reqAmount
			amount = math.Round(exchangeData.Amount*100) / 100
		} else {
			rate = reqAmount / exchangeData.Amount
			amount = exchangeData.Amount
		}

		return models.ExchangeRate{
			Exchange: "Bitrush",
			Rate:     math.Round(rate*100) / 100,
			Amount:   amount,
			Link:     "https://bitrush.nl/nl?r=bitdeal",
			Reviews:  7,
			Stars:    4,
		}
	case *models.BitrushSell:
		reqAmount, _ := strconv.ParseFloat(request.Amount, 64)
		if request.Currency == "btc" {
			rate = exchangeData.Amount / reqAmount
			amount = math.Round(exchangeData.Amount*100) / 100
		} else {
			rate = reqAmount / exchangeData.Amount
			amount = exchangeData.Amount
		}

		return models.ExchangeRate{
			Exchange: "Bitrush",
			Rate:     math.Round(rate*100) / 100,
			Amount:   amount,
			Link:     "https://bitrush.nl/nl?r=bitdeal",
			Reviews:  7,
			Stars:    4,
		}

	case *models.KnakenBuy:
		reqAmount, _ := strconv.ParseFloat(request.Amount, 64)
		respAmount, _ := strconv.ParseFloat(exchangeData.Price, 64)

		if request.Currency == "btc" {
			rate = respAmount / reqAmount
			amount = math.Round(respAmount*100) / 100
		} else {
			rate = reqAmount / respAmount
			amount = respAmount
		}

		return models.ExchangeRate{
			Exchange: "Knaken",
			Rate:     math.Round(rate*100) / 100,
			Amount:   amount,
			Link:     "knaken.eu/?ref=5da47c2e7823c",
			Reviews:  12,
			Stars:    4,
		}
	case *models.KnakenSell:
		reqAmount, _ := strconv.ParseFloat(request.Amount, 64)
		respRate, _ := strconv.ParseFloat(exchangeData.BTC.SellPrice, 64)

		if request.Currency == "btc" {
			rate = respRate
			amount = math.Round(respRate*reqAmount*100) / 100
		} else {
			rate = respRate
			amount = math.Round(reqAmount/respRate*100000000) / 100000000
		}

		return models.ExchangeRate{
			Exchange: "Knaken",
			Rate:     math.Round(rate*100) / 100,
			Amount:   amount,
			Link:     "knaken.eu/?ref=5da47c2e7823c",
			Reviews:  12,
			Stars:    4,
		}
	}

	return models.ExchangeRate{}
}
