package handlers

import (
	"math"

	"bitdeal.nl/models"
)

func wrap(exchangeRate models.ExchangeRate, bestRate float64, mostReviews int) map[string]interface{} {
	return map[string]interface{}{
		"Exchange":    exchangeRate,
		"BestRate":    bestRate,
		"MostReviews": mostReviews,
	}
}

func bestRateToBTC(bestRate float64) float64 {
	return math.Round(420/bestRate*100000000) / 100000000
}

func trimCaptcha(s string) string {
	m := 0
	for i := range s {
		if m >= 22 {
			return s[i:]
		}
		m++
	}
	return s[:0]
}

func printErrorParam(s string) string {
	switch s {
	case "passwordNumber":
		return "Password didn't include a number"
	case "passwordUpper":
		return "Password didn't include an uppercase letter"
	case "passwordCount":
		return "Password wasn't 8 characters or longer"
	case "captcha":
		return "Wrong captcha"
	case "usernameChars":
		return "The username included an illegal character"
	case "usernameLength":
		return "The username wasn't 3 characters or longer"
	case "expired":
		return "The captcha expired"
	}

	return "Error occured, try again"
}
