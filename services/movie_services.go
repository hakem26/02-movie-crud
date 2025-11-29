package services

import (
	"errors"
	"example/moviecrud/models"
	"math/rand"
	"strconv"
)

func CreateMovie(movie models.Movie) (models.Movie, error) {
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	models.Movies = append(models.Movies, movie)
	return movie, nil
}

func GetAllMovies() []models.Movie {
	return models.Movies
}

func GetMovieByID(id string) (models.Movie, error) {
	for _, m := range models.Movies {
		if m.ID == id {
			return m, nil
		}
	}
	return models.Movie{}, errors.New("movie not found")
}

func UpdateMovie(id string, newData models.Movie) (models.Movie, error) {
	for i, m := range models.Movies {
		if m.ID == id {
			newData.ID = id
			models.Movies[i] = newData
			return newData, nil
		}
	}
	return models.Movie{}, errors.New("movie not found")
}

func DeleteMovie(id string) error {
	for i, m := range models.Movies {
		if m.ID == id {
			models.Movies = append(models.Movies[:i], models.Movies[i+1:]...)
			return nil
		}
	}
	return errors.New("movie not found")
}