package models

/*
Responsetime ...
*/
type Responsetime struct {
	ID          int64   `json:"id"`
	Requesttime int64   `json:"requesttime"`
	Created     int64   `json:"created"`
	Difference  float64 `json:"difference"`
}

/*
Responsetimes ...
*/
type Responsetimes struct {
	Average       int64          `json:"average"`
	Variance      float64        `json:"variance"`
	Sigma         float64        `json:"sigma"`
	Responsetimes []Responsetime `json:"responsetimes"`
}
