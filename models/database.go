package models

/*
Responsetime ...
*/
type Responsetime struct {
	ID          int64   `json:"id"`
	Requesttime int64   `json:"requesttime"`
	Created     int64   `json:"created"`
	Difference  float64 `json:"difference"`
	Zscore      float64 `json:"zscore"`
	CDF         float64 `json:"cdf"`
}

/*
Responsetimes ...
*/
type Responsetimes struct {
	Average       int64          `json:"average"`
	Variance      float64        `json:"variance"`
	Sigma         float64        `json:"sigma"`
	Count         int64          `json:"count"`
	Responsetimes []Responsetime `json:"responsetimes"`
}

/*
ResponseSunshade ...
*/
type ResponseSunshade struct {
	ID     int64 `json:"id"`
	Length int64 `json:"length"`
}
