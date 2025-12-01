package repository

import "example/moviecrud/models"

type MovieRepository interface {
	Create(m *models.Movie) error
	FindAll() ([]*models.Movie, error)
	FindByID(id string) (*models.Movie, error) 
	FindByDirector(director string) (*models.Director, error)
	Update(id string, m *models.Movie) (*models.Movie, error) 
	Delete(id string) error
}