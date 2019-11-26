package etl

import (
	"math"
	"strconv"

	"bitdeal.nl/models"
)

var (
	// Bitvavo ...
	Bitvavo = models.ExchangeInfo{
		Name:    "Bitvavo",
		Link:    "https://bitvavo.com/?a=C23D4F6D97",
		Reviews: 1895,
		Stars:   5,
		Broker:  false,
	}

	// Litebit ...
	Litebit = models.ExchangeInfo{
		Name:    "Litebit",
		Link:    "https://www.litebit.eu/?referrer=40203",
		Reviews: 1695,
		Stars:   5,
		Broker:  true,
	}

	// Bitonic ...
	Bitonic = models.ExchangeInfo{
		Name:    "Bitonic",
		Link:    "https://bitonic.nl/?partner=777",
		Reviews: 41,
		Stars:   4,
		Broker:  true,
	}

	// Btcdirect ...
	Btcdirect = models.ExchangeInfo{
		Name:    "Btcdirect",
		Link:    "https://btcdirect.eu/nl-nl?partnerId=319",
		Reviews: 5485,
		Stars:   5,
		Broker:  true,
	}

	// Bitrush ...
	Bitrush = models.ExchangeInfo{
		Name:    "Bitrush",
		Link:    "https://bitrush.nl/nl?r=bitdeal",
		Reviews: 7,
		Stars:   4,
		Broker:  true,
	}

	// Knaken ...
	Knaken = models.ExchangeInfo{
		Name:    "Knaken",
		Link:    "https://knaken.eu/?ref=5da47c2e7823c",
		Reviews: 12,
		Stars:   4,
		Broker:  true,
	}
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
			Exchange: Bitvavo.Name,
			Rate:     math.Round(rate*100) / 100,
			Amount:   amount,
			Link:     Bitvavo.Link,
			Reviews:  Bitvavo.Reviews,
			Stars:    Bitvavo.Stars,
			Status:   models.OK,
			Broker:   Bitvavo.Broker,
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
			Exchange: Litebit.Name,
			Rate:     math.Round(rate*100) / 100,
			Amount:   amount,
			Link:     Litebit.Link,
			Reviews:  Litebit.Reviews,
			Stars:    Litebit.Stars,
			Status:   models.OK,
			Broker:   Litebit.Broker,
		}
	case *models.BitonicBuy:
		if request.Currency == "btc" {
			amount = math.Round(exchangeData.Eur*100) / 100
		} else {
			amount = exchangeData.Btc
		}

		return models.ExchangeRate{
			Exchange: Bitonic.Name,
			Rate:     math.Round(exchangeData.Price*100) / 100,
			Amount:   amount,
			Link:     Bitonic.Link,
			Reviews:  Bitonic.Reviews,
			Stars:    Bitonic.Stars,
			Status:   models.OK,
			Broker:   Bitonic.Broker,
		}
	case *models.BitonicSell:
		if request.Currency == "btc" {
			amount = math.Round(exchangeData.Eur*100) / 100
		} else {
			amount = exchangeData.Btc
		}

		return models.ExchangeRate{
			Exchange: Bitonic.Name,
			Rate:     math.Round(exchangeData.Price*100) / 100,
			Amount:   amount,
			Link:     Bitonic.Link,
			Reviews:  Bitonic.Reviews,
			Stars:    Bitonic.Stars,
			Status:   models.OK,
			Broker:   Bitonic.Broker,
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
			Exchange: Btcdirect.Name,
			Rate:     math.Round(rate*100) / 100,
			Amount:   amount,
			Link:     Btcdirect.Link,
			Reviews:  Btcdirect.Reviews,
			Stars:    Btcdirect.Stars,
			Status:   models.OK,
			Broker:   Btcdirect.Broker,
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
			Exchange: Bitrush.Name,
			Rate:     math.Round(rate*100) / 100,
			Amount:   amount,
			Link:     Bitrush.Link,
			Reviews:  Bitrush.Reviews,
			Stars:    Bitrush.Stars,
			Status:   models.OK,
			Broker:   Bitrush.Broker,
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
			Exchange: Bitrush.Name,
			Rate:     math.Round(rate*100) / 100,
			Amount:   amount,
			Link:     Bitrush.Link,
			Reviews:  Bitrush.Reviews,
			Stars:    Bitrush.Stars,
			Status:   models.OK,
			Broker:   Bitrush.Broker,
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
			Exchange: Knaken.Name,
			Rate:     math.Round(rate*100) / 100,
			Amount:   amount,
			Link:     Knaken.Link,
			Reviews:  Knaken.Reviews,
			Stars:    Knaken.Stars,
			Status:   models.OK,
			Broker:   Knaken.Broker,
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
			Link:     Knaken.Link,
			Reviews:  Knaken.Reviews,
			Stars:    Knaken.Stars,
			Status:   models.OK,
			Broker:   Knaken.Broker,
		}
	case *models.Coinmerce:
		var respRate float64
		reqAmount, _ := strconv.ParseFloat(request.Amount, 64)

		if request.Type == "buy" {
			respRate, _ = strconv.ParseFloat(exchangeData.Prices.BTC.Buy, 64)
			rate = respRate
		} else {
			respRate, _ = strconv.ParseFloat(exchangeData.Prices.BTC.Sell, 64)
			rate = respRate
		}

		if request.Currency == "btc" {
			amount = math.Round(respRate*reqAmount*100) / 100
		} else {
			amount = math.Round(reqAmount/respRate*100000000) / 100000000
		}

		return models.ExchangeRate{
			Exchange: "Coinmerce",
			Rate:     math.Round(rate*100) / 100,
			Amount:   amount,
			Link:     "https://coinmerce.io/r/9WZT8VCHMo",
			Reviews:  107,
			Stars:    5,
			Status:   models.OK,
			Broker:   true,
		}
	case *models.GetPricesError:
		return returnError(exchangeData)
	}

	return models.ExchangeRate{}
}

func returnError(exchangeData *models.GetPricesError) models.ExchangeRate {
	switch exchangeData.Exchange {
	case "Bitvavo":
		return models.ExchangeRate{
			Exchange: Bitvavo.Name,
			Rate:     1,
			Amount:   1,
			Link:     Bitvavo.Link,
			Reviews:  Bitvavo.Reviews,
			Stars:    Bitvavo.Stars,
			Status:   exchangeData.Error,
			Broker:   Bitvavo.Broker,
		}
	case "Litebit":
		return models.ExchangeRate{
			Exchange: Litebit.Name,
			Rate:     1,
			Amount:   1,
			Link:     Litebit.Link,
			Reviews:  Litebit.Reviews,
			Stars:    Litebit.Stars,
			Status:   exchangeData.Error,
			Broker:   Litebit.Broker,
		}
	case "Bitonic":
		return models.ExchangeRate{
			Exchange: Bitonic.Name,
			Rate:     1,
			Amount:   1,
			Link:     Bitonic.Link,
			Reviews:  Bitonic.Reviews,
			Stars:    Bitonic.Stars,
			Status:   exchangeData.Error,
			Broker:   Bitonic.Broker,
		}
	case "Btcdirect":
		return models.ExchangeRate{
			Exchange: Btcdirect.Name,
			Rate:     1,
			Amount:   1,
			Link:     Btcdirect.Link,
			Reviews:  Btcdirect.Reviews,
			Stars:    Btcdirect.Stars,
			Status:   exchangeData.Error,
			Broker:   Btcdirect.Broker,
		}
	case "Bitrush":
		return models.ExchangeRate{
			Exchange: Btcdirect.Name,
			Rate:     1,
			Amount:   1,
			Link:     Btcdirect.Link,
			Reviews:  Btcdirect.Reviews,
			Stars:    Btcdirect.Stars,
			Status:   exchangeData.Error,
			Broker:   Btcdirect.Broker,
		}
	case "Knaken":
		return models.ExchangeRate{
			Exchange: Knaken.Name,
			Rate:     1,
			Amount:   1,
			Link:     Knaken.Link,
			Reviews:  Knaken.Reviews,
			Stars:    Knaken.Stars,
			Status:   exchangeData.Error,
			Broker:   Knaken.Broker,
		}
	}

	return models.ExchangeRate{}
}
