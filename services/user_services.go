package services

import (
	"errors"
	"example/moviecrud/models"
	"example/moviecrud/repository"
	"example/moviecrud/utils"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) CreateUser(u *models.User) (*models.User, error) {
	// basic validation
	if u.Email == "" || u.Password == "" {
		return nil, errors.New("email and password required")
	}

	// check existing email
	if existing, _ := s.repo.FindByEmail(u.Email); existing != nil {
		return nil, errors.New("email already used")
	}

	hashed, err := utils.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	u.Password = hashed

	if err := s.repo.Create(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *UserService) GetAll() ([]*models.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) GetByID(id string) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) Update(id string, u *models.User) (*models.User, error) {
	// اگر پسورد داده شده، هش کن
	if u.Password != "" && len(u.Password) < 60 { //  // 60 ≈ طول bcrypt hash
		hashed, err := utils.HashPassword(u.Password)
		if err != nil {
			return nil, err
		}
		u.Password = hashed
	}

	updatedUser, err := s.repo.Update(id, u)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s *UserService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *UserService) GetByUserID(userID string) (*models.User, error) {
	return s.repo.FindByUserID(userID) // باید توی repo هم اضافه کنی (یه خط)
}
