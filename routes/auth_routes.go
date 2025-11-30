package routes

import (
	"example/moviecrud/controllers"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(r *mux.Router, ac *controllers.AuthController) {
	auth := r.PathPrefix("/auth").Subrouter()

	auth.HandleFunc("/register", ac.Register).Methods("POST")
	auth.HandleFunc("/login", ac.Login).Methods("POST")
	auth.HandleFunc("/refresh", ac.RefreshToken).Methods("POST")
}