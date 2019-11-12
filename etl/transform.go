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

		if request.Currency == "btc" {
			amount = math.Round(request.Amount*rate*100) / 100
		} else {
			amount = math.Round(request.Amount/rate*100000000) / 100000000
		}

		return models.ExchangeRate{
			Exchange: "Bitvavo",
			Rate:     math.Round(rate*100) / 100,
			Amount:   amount,
		}
	case *models.Litebit:
		if request.Type == "buy" {
			rate, _ = strconv.ParseFloat(exchangeData.Result.Buy, 32)
		} else {
			rate, _ = strconv.ParseFloat(exchangeData.Result.Sell, 32)
		}

		if request.Currency == "btc" {
			amount = math.Round(request.Amount*rate*100) / 100
		} else {
			amount = math.Round(request.Amount/rate*100000000) / 100000000
		}

		return models.ExchangeRate{
			Exchange: "Litebit",
			Rate:     math.Round(rate*100) / 100,
			Amount:   amount,
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
		}
	case *models.Btcdirect:
		if request.Type == "buy" {
			rate, _ = strconv.ParseFloat(exchangeData.ExchangeRates.Sell.Btc, 32)
		} else {
			rate, _ = strconv.ParseFloat(exchangeData.ExchangeRates.Buy.Btc, 32)
		}

		if request.Currency == "btc" {
			amount = math.Round(request.Amount*rate*100) / 100
		} else {
			amount = math.Round(request.Amount/rate*100000000) / 100000000
		}

		return models.ExchangeRate{
			Exchange: "Btcdirect",
			Rate:     math.Round(rate*100) / 100,
			Amount:   amount,
		}
	case *models.BitrushBuy:
		if request.Currency == "btc" {
			rate = exchangeData.Amount / request.Amount
			amount = math.Round(exchangeData.Amount*100) / 100
		} else {
			rate = request.Amount / exchangeData.Amount
			amount = exchangeData.Amount
		}

		return models.ExchangeRate{
			Exchange: "Bitrush",
			Rate:     math.Round(rate*100) / 100,
			Amount:   amount,
		}
	case *models.BitrushSell:
		if request.Currency == "btc" {
			rate = exchangeData.Amount / request.Amount
			amount = math.Round(exchangeData.Amount*100) / 100
		} else {
			rate = request.Amount / exchangeData.Amount
			amount = exchangeData.Amount
		}

		return models.ExchangeRate{
			Exchange: "Bitrush",
			Rate:     math.Round(rate*100) / 100,
			Amount:   amount,
		}
	}

	return models.ExchangeRate{}
}
