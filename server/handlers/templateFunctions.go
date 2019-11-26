package handlers

import (
	"math"

	"bitdeal.nl/models"
)

func wrap(exchangeRate models.ExchangeRate, bestRate float64, mostReviews int, iphone bool) map[string]interface{} {
	return map[string]interface{}{
		"Exchange":    exchangeRate,
		"BestRate":    bestRate,
		"MostReviews": mostReviews,
		"Iphone":      iphone,
	}
}

func bestRateToBTC(bestRate float64) float64 {
	return math.Round(420/bestRate*100000000) / 100000000
}

func stars(stars int) string {
	if stars == 4 {
		return "★★★★"
	}
	return "★★★★★"
}
