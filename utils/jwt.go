package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	AccessSecret  = []byte("s3cr3t-plz-change-in-real-project-2025")
	refreshSecret = []byte("another-super-secret-refresh-2025")
	AccessTTL     = 15 * time.Minute
	RefreshTTL    = 7 * 24 * time.Hour
)

type Claims struct {
	UserID    string `json:"user_id"`
	Email     string `json:"email"`
	LevelName string `json:"level_name"`
	jwt.RegisteredClaims
}

func GenerateTokens(userID, email, levelName string) (string, string, error) {
	access := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserID:    userID,
		Email:     email,
		LevelName: levelName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.New().String(),
		},
	})

	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshTTL)),
		Subject:   userID,
		ID:        uuid.New().String(),
	})

	a, err := access.SignedString(AccessSecret)
	if err != nil { return "", "", err }
	r, err := refresh.SignedString(refreshSecret)
	if err != nil { return "", "", err }

	return a, r, nil
}