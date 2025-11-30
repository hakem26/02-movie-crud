package services

import (
	"errors"

	"example/moviecrud/dto"
	"example/moviecrud/models"
	"example/moviecrud/repository"
	"example/moviecrud/utils"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthService struct {
	userRepo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{userRepo: repo}
}

func (s *AuthService) Register(req *dto.RegisterRequest) (*models.PublicUser, error) {
	// چک کردن وجود ایمیل (دوبل چک با ایندکس)
	if existing, _ := s.userRepo.FindByEmail(req.Email); existing != nil {
		return nil, errors.New("email already registered")
	}

	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		UserID:   uuid.New().String(),
		FullName: req.FullName,
		Email:    req.Email,
		Password: hashed,
		Level: models.UserLevel{
			LevelID:   req.Level.LevelID,   // اینا دقیقاً یکسان هستن
			LevelName: req.Level.LevelName,
		},
	}

	if err := s.userRepo.Create(user); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, errors.New("email already registered")
		}
		return nil, err
	}

	return &models.PublicUser{
		ID:       user.UserID,
		FullName: user.FullName,
		Email:    user.Email,
		Level:    user.Level,
	}, nil
}

func (s *AuthService) Login(email, password string) (*dto.TokenResponse, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil || user == nil || !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	access, refresh, err := utils.GenerateTokens(user.UserID, user.Email, user.Level.LevelName)
	if err != nil {
		return nil, err
	}

	return &dto.TokenResponse{
		AccessToken:  access,
		RefreshToken: refresh,
		TokenType:    "Bearer",
		ExpiresIn:    900,
	}, nil
}