package controllers

import (
	"encoding/json"
	"example/moviecrud/models"
	"example/moviecrud/services"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services.GetAllMovies())
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	movie, err := services.GetMovieByID(id)
	if err == nil {
		json.NewEncoder(w).Encode(movie)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var movie models.Movie
	json.NewDecoder(r.Body).Decode(&movie)

	created, err := services.CreateMovie(movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(created)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var data models.Movie
	json.NewDecoder(r.Body).Decode(&data)
	updated, err := services.UpdateMovie(id, data)
	if err == nil {
		json.NewEncoder(w).Encode(updated)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	err := services.DeleteMovie(id)
	if err == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
