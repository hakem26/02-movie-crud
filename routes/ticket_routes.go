package routes

import (
	"example/moviecrud/controllers"

	"github.com/gorilla/mux"
)

func RegisterTicketRoutes(router *mux.Router) {
	router.HandleFunc("/tickets", controllers.GetTickets).Methods("GET")
	router.HandleFunc("/tickets/{id}", controllers.GetTicket).Methods("GET")
	router.HandleFunc("/tickets", controllers.CreateTicket).Methods("POST")
	router.HandleFunc("/tickets/{id}", controllers.UpdateTicket).Methods("PUT")
	router.HandleFunc("/tickets/{id}", controllers.DeleteTicket).Methods("DELETE")
}