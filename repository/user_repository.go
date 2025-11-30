package repository

import "example/moviecrud/models"

type UserRepository interface {
	Create(u *models.User) error
	FindAll() ([]*models.User, error)
	FindByID(id string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByUserID(userID string) (*models.User, error)
	Update(id string, u *models.User) (*models.User, error)  // این دو مقدار برمی‌گردونه
	Delete(id string) error
}