package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserLevel struct {
	LevelID   string `json:"level_id" bson:"level_id" validate:"required"`
	LevelName string `json:"level_name" bson:"level_name" validate:"required"`
}

type User struct {
	ID       primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	UserID   string             `json:"id" bson:"user_id"`
	FullName string             `json:"fullname" bson:"fullname" validate:"required,min=3"`
	Email    string             `json:"email" bson:"email" validate:"required,email"`
	Password string             `json:"password,omitempty" bson:"password" validate:"required,min=8"`
	Level    UserLevel          `json:"level" bson:"level" validate:"required"`
}

// برای نمایش عمومی (بدون پسورد)
type PublicUser struct {
	ID       string    `json:"id"`
	FullName string    `json:"fullname"`
	Email    string    `json:"email"`
	Level    UserLevel `json:"level"`
}