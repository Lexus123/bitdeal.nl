package models

/*
HomepageData ...
*/
type HomepageData struct {
	Title          string
	ExchangePrices GetPricesResponse
	URL            string
}

/*
StatsData ...
*/
type StatsData struct {
	Title string
	Stats interface{}
	URL   string
}
