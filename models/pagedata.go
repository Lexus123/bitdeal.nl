package models

/*
HomepageData ...
*/
type HomepageData struct {
	Title          string
	ExchangePrices GetPricesResponse
}

/*
StatsData ...
*/
type StatsData struct {
	Title string
	Stats interface{}
}
