package etl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"bitdeal.nl/models"
)

/*
Exchanges ...
*/
var Exchanges = []models.Exchange{
	{
		Name:     "Bitvavo",
		URL:      "https://api.bitvavo.com/",
		Method:   "GET",
		Type:     "both",
		Currency: "both",
		Endpoint: "v2/BTC-EUR/book",
	},
	{
		Name:     "Litebit",
		URL:      "https://api.litebit.eu/",
		Method:   "GET",
		Type:     "both",
		Currency: "both",
		Endpoint: "market/btc",
	},
	{
		Name:     "BitonicBuy",
		URL:      "https://bitonic.nl/",
		Method:   "GET",
		Type:     "buy",
		Currency: "eur",
		Endpoint: "api/buy?eur=",
	},
	{
		Name:     "BitonicBuy",
		URL:      "https://bitonic.nl/",
		Method:   "GET",
		Type:     "buy",
		Currency: "btc",
		Endpoint: "api/buy?btc=",
	},
	{
		Name:     "BitonicSell",
		URL:      "https://bitonic.nl/",
		Method:   "GET",
		Type:     "sell",
		Currency: "eur",
		Endpoint: "api/sell?eur=",
	},
	{
		Name:     "BitonicSell",
		URL:      "https://bitonic.nl/",
		Method:   "GET",
		Type:     "sell",
		Currency: "btc",
		Endpoint: "api/sell?btc=",
	},
	{
		Name:     "Btcdirect",
		URL:      "https://my.btcdirect.eu/",
		Method:   "GET",
		Type:     "both",
		Currency: "both",
		Endpoint: "config",
	},
	{
		Name:     "KnakenBuy",
		URL:      "https://knaken.eu/",
		Method:   "GET",
		Type:     "buy",
		Currency: "eur",
		Endpoint: "api/price/calculate-buy-price-cryptocurrency/BTC/",
	},
	{
		Name:     "KnakenBuy",
		URL:      "https://knaken.eu/",
		Method:   "GET",
		Type:     "buy",
		Currency: "btc",
		Endpoint: "api/price/calculate-buy-price-eur/BTC/",
	},
	{
		Name:     "KnakenSell",
		URL:      "https://knaken.eu/",
		Method:   "GET",
		Type:     "sell",
		Currency: "both",
		Endpoint: "api/get-prices",
	},
	// {
	// 	Name:     "Coinmerce",
	// 	URL:      "https://cmapi.nl/",
	// 	Method:   "GET",
	// 	Type:     "both",
	// 	Currency: "both",
	// 	Endpoint: "pricesall",
	// },
	{
		Name:     "Satos",
		URL:      "https://api.satos.eu/",
		Method:   "GET",
		Type:     "buy",
		Currency: "eur",
		Endpoint: "ticker/eur/btc/buy/fiat/",
	},
	{
		Name:     "Satos",
		URL:      "https://api.satos.eu/",
		Method:   "GET",
		Type:     "buy",
		Currency: "btc",
		Endpoint: "ticker/eur/btc/buy/coin/",
	},
	{
		Name:     "Satos",
		URL:      "https://api.satos.eu/",
		Method:   "GET",
		Type:     "sell",
		Currency: "eur",
		Endpoint: "ticker/eur/btc/sell/fiat/",
	},
	{
		Name:     "Satos",
		URL:      "https://api.satos.eu/",
		Method:   "GET",
		Type:     "sell",
		Currency: "btc",
		Endpoint: "ticker/eur/btc/sell/coin/",
	},
}

func structSelector(body []byte, s interface{}) (interface{}, error) {
	err := json.Unmarshal(body, &s)
	return s, err
}

func toJSON(body []byte, exchange models.Exchange) (interface{}, error) {
	switch exchange.Name {
	case "Bitvavo":
		return structSelector(body, new(models.Bitvavo))
	case "Litebit":
		return structSelector(body, new(models.Litebit))
	case "BitonicBuy":
		return structSelector(body, new(models.BitonicBuy))
	case "BitonicSell":
		return structSelector(body, new(models.BitonicSell))
	case "Btcdirect":
		return structSelector(body, new(models.Btcdirect))
	case "BitrushBuy":
		return structSelector(body, new(models.BitrushBuy))
	case "BitrushSell":
		return structSelector(body, new(models.BitrushSell))
	case "KnakenBuy":
		return structSelector(body, new(models.KnakenBuy))
	case "KnakenSell":
		return structSelector(body, new(models.KnakenSell))
	case "Coinmerce":
		return structSelector(body, new(models.Coinmerce))
	case "Satos":
		return structSelector(body, new(models.Satos))
	}

	return nil, nil
}

func fetchPrices(request models.GetPricesData, exchange models.Exchange, dataChannel chan interface{}) {
	var resp *http.Response
	var err error
	var s interface{}

	switch exchange.Name {
	case "Bitvavo":
		resp, err = http.Get(exchange.URL + exchange.Endpoint)
		break
	case "Litebit":
		resp, err = http.Get(exchange.URL + exchange.Endpoint)
		break
	case "BitonicBuy":
		resp, err = http.Get(exchange.URL + exchange.Endpoint + request.Amount)
		break
	case "BitonicSell":
		resp, err = http.Get(exchange.URL + exchange.Endpoint + request.Amount)
		break
	case "Btcdirect":
		resp, err = http.Get(exchange.URL + exchange.Endpoint)
		break
	case "BitrushBuy":
		resp, err = http.Get(exchange.URL + exchange.Endpoint + request.Amount)
		break
	case "BitrushSell":
		resp, err = http.Get(exchange.URL + exchange.Endpoint + request.Amount)
		break
	case "KnakenBuy":
		resp, err = http.Get(exchange.URL + exchange.Endpoint + request.Amount + "/N/N/0")
		break
	case "KnakenSell":
		resp, err = http.Get(exchange.URL + exchange.Endpoint)
		break
	case "Coinmerce":
		resp, err = http.Get(exchange.URL + exchange.Endpoint)
		break
	case "Satos":
		resp, err = http.Get(exchange.URL + exchange.Endpoint + request.Amount)
		break
	}

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		respData, err := ioutil.ReadAll(resp.Body)
		s, err = toJSON([]byte(respData), exchange)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		s = models.GetPricesError{
			Exchange: exchange.Name,
			Error:    models.GeneralError,
		}
	}

	dataChannel <- s
}

/*
DoRequests ...
*/
func DoRequests(request models.GetPricesData, dataChannel chan interface{}) {
	for _, exchange := range Exchanges {
		if exchange.Type == request.Type || exchange.Type == "both" {
			if exchange.Currency == request.Currency || exchange.Currency == "both" {
				go fetchPrices(request, exchange, dataChannel)
			}
		}
	}
}
