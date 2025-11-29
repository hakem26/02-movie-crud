package models

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

var Movies []Movie
