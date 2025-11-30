package controllers

import (
	"encoding/json"
	"example/moviecrud/models"
	"example/moviecrud/services"
	"net/http"

	"github.com/gorilla/mux"
)

func GetTickets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(services.GetAllTickets())
}

func GetTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	ticket, err := services.GetTicketByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(ticket)
}

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var ticket models.Ticket
	json.NewDecoder(r.Body).Decode(&ticket)
	created, err := services.CreateTicket(ticket)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(created)
}

func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	var data models.Ticket
	json.NewDecoder(r.Body).Decode(&data)
	updated, err := services.UpdateTicket(id, data)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updated)
}

func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	err := services.DeleteTicket(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "ticket deleted"})
}

