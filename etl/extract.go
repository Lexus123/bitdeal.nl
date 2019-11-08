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
		URL:      "https://api.bitvavo.com/v2/BTC-EUR/book",
		Method:   "GET",
		Type:     "both",
		Currency: "both",
	},
	{
		Name:     "Litebit",
		URL:      "https://api.litebit.eu/market/btc",
		Method:   "GET",
		Type:     "both",
		Currency: "both",
	},
}

func structSelector(body []byte, s interface{}) (interface{}, error) {
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

func toJSON(body []byte, exchange models.Exchange) (interface{}, error) {
	switch exchange.Name {
	case "Bitvavo":
		return structSelector(body, new(models.Bitvavo))
	case "Litebit":
		return structSelector(body, new(models.Litebit))
	}

	return nil, nil
}

func fetchPrices(exchange models.Exchange, dataChannel chan interface{}) {
	var resp *http.Response
	var err error

	if exchange.Method == "GET" {
		resp, err = http.Get(exchange.URL)
	} else {
		resp, err = http.Post(exchange.URL, "application/json; charset=utf-8", nil)
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
				go fetchPrices(exchange, dataChannel)
			}
		}
	}
}
