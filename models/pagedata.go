package models

/*
HomepageData ...
*/
type HomepageData struct {
	Title          string
	Iphone         bool
	ExchangePrices GetPricesResponse
	URL            string
}

/*
StatsData ...
*/
type StatsData struct {
	Title string
	URL   string
}
