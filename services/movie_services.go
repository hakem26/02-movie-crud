package services

import "example/moviecrud/repository"

type MovieService struct {
	repo repository.MovieRepository
}

func NewMovieService(r repository.MovieRepository) *MovieService {
	return &MovieService{repo: r}
}

