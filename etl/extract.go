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
		Name:     "BitrushBuy",
		URL:      "https://api.bitrush.nl/",
		Method:   "GET",
		Type:     "buy",
		Currency: "eur",
		Endpoint: "v1/products/EUR-XBT/ask?currency=EUR&amount=",
	},
	{
		Name:     "BitrushBuy",
		URL:      "https://api.bitrush.nl/",
		Method:   "GET",
		Type:     "buy",
		Currency: "btc",
		Endpoint: "v1/products/EUR-XBT/ask?currency=XBT&amount=",
	},
	{
		Name:     "BitrushSell",
		URL:      "https://api.bitrush.nl/",
		Method:   "GET",
		Type:     "sell",
		Currency: "eur",
		Endpoint: "v1/products/EUR-XBT/bid?currency=EUR&amount=",
	},
	{
		Name:     "BitrushSell",
		URL:      "https://api.bitrush.nl/",
		Method:   "GET",
		Type:     "sell",
		Currency: "btc",
		Endpoint: "v1/products/EUR-XBT/bid?currency=XBT&amount=",
	},
}

func structSelector(body []byte, s interface{}) (interface{}, error) {
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println(err)
	}
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
	}

	return nil, nil
}

func fetchPrices(request models.GetPricesData, exchange models.Exchange, dataChannel chan interface{}) {
	var resp *http.Response
	var err error

	switch exchange.Name {
	case "Bitvavo":
		resp, err = http.Get(exchange.URL + exchange.Endpoint)
		break
	case "Litebit":
		resp, err = http.Get(exchange.URL + exchange.Endpoint)
		break
	case "BitonicBuy":
		amount := fmt.Sprintf("%f", request.Amount)
		resp, err = http.Get(exchange.URL + exchange.Endpoint + amount)
		break
	case "BitonicSell":
		amount := fmt.Sprintf("%f", request.Amount)
		resp, err = http.Get(exchange.URL + exchange.Endpoint + amount)
		break
	case "Btcdirect":
		resp, err = http.Get(exchange.URL + exchange.Endpoint)
		break
	case "BitrushBuy":
		amount := fmt.Sprintf("%f", request.Amount)
		resp, err = http.Get(exchange.URL + exchange.Endpoint + amount)
		break
	case "BitrushSell":
		amount := fmt.Sprintf("%f", request.Amount)
		resp, err = http.Get(exchange.URL + exchange.Endpoint + amount)
		break
	}

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	s, err := toJSON([]byte(respData), exchange)
	if err != nil {
		fmt.Println(err)
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
