package models

type Ticket struct {
	ID 	  string  `json:"id"`
	MovieID string  `json:"movie_id"`
	Price   float64 `json:"price"`
	Seat    string  `json:"seat"`
}

var Tickets []Ticket