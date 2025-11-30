package routes

import (
	"example/moviecrud/controllers"
	"example/moviecrud/middleware"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router, uc *controllers.UserController) {
	user := r.PathPrefix("/users").Subrouter()

	// همه این روت‌ها نیاز به لاگین دارن
	user.Use(middleware.Auth(false))

	user.HandleFunc("", uc.GetUsers).Methods("GET")
	user.HandleFunc("/me", uc.GetMe).Methods("GET")
	user.HandleFunc("/me", uc.UpdateMe).Methods("PUT")
	user.HandleFunc("/{id}", uc.GetUser).Methods("GET")

	// این دو تا فقط ادمین
	admin := r.PathPrefix("/users").Subrouter()
	admin.Use(middleware.Auth(true)) // فقط ادمین
	admin.HandleFunc("/{id}", uc.UpdateUser).Methods("PUT")
	admin.HandleFunc("/{id}", uc.DeleteUser).Methods("DELETE")
}
