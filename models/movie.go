package models

type Director struct {
	FirstName string `json:"firstname" bson:"firstname" validate:"required"`
	LastName  string `json:"lastname" bson:"lastname" validate:"required"`
}

type Movie struct {
	ID       string    `json:"-" bson:"_id,omitempty"`
	Isbn     string    `json:"isbn" bson:"isbn" validate:"required"`
	Title    string    `json:"title" bson:"title" validate:"required,min=3"`
	Director *Director `json:"director" bson:"director" validate:"required"`
}