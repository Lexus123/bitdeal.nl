package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"

	"bitdeal.nl/etl"
	"bitdeal.nl/models"
)

/*
GetPrices ...
*/
func GetPrices(w http.ResponseWriter, r *http.Request) {
	// Read JSON body of request and check for errors. For example: buy, eur, 100
	bodyAsByte, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal the body and check for errors
	var message models.GetPricesData
	err = json.Unmarshal(bodyAsByte, &message)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Send message to the extractor to get data from brokers
	dataChannel := make(chan interface{}, 5)
	defer close(dataChannel)
	go etl.DoRequests(message, dataChannel)

	// Prepare the response message
	getPricesResponse := models.GetPricesResponse{
		Type:     message.Type,
		Currency: message.Currency,
	}

	// For each exchange, transform the data and put it in the response
	for i := 0; i < 5; i++ {
		getPricesResponse.ExchangeRates = append(getPricesResponse.ExchangeRates, etl.Transform(message, <-dataChannel))
	}

	// Sort the prices according to the request
	if getPricesResponse.Type == "buy" {
		sort.Slice(getPricesResponse.ExchangeRates[:], func(i, j int) bool {
			return getPricesResponse.ExchangeRates[i].Rate < getPricesResponse.ExchangeRates[j].Rate
		})
	} else {
		sort.Slice(getPricesResponse.ExchangeRates[:], func(i, j int) bool {
			return getPricesResponse.ExchangeRates[i].Rate > getPricesResponse.ExchangeRates[j].Rate
		})
	}

	getPricesResponse.BestRate = getPricesResponse.ExchangeRates[0].Rate
	getPricesResponse.BestAmount = getPricesResponse.ExchangeRates[0].Amount
	getPricesResponse.MostReviews = 5485

	output, err := json.Marshal(getPricesResponse)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
