package models

type UserLevel struct {
	LevelID string `json:"level_id"`
	LevelName string `json:"level_name"`
}

type User struct {
	ID string `json:"id"`
	FullName string `json:"fullname"`
	Email string `json:"email"`
	Password string `json:"password"`
	Level UserLevel `json:"level"`
}

var Users []User