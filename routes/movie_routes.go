package routes

import (
	"example/moviecrud/controllers"
	"github.com/gorilla/mux"
)

func RegisterMovieRoutes(router *mux.Router) {
	router.HandleFunc("/movies", controllers.GetMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", controllers.GetMovie).Methods("GET")
	router.HandleFunc("/movies", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", controllers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", controllers.DeleteMovie).Methods("DELETE")
}
