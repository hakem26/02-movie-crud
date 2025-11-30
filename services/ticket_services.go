package services

import (
	"errors"
	"example/moviecrud/models"
	"math/rand"
	"strconv"
)

func GetAllTickets() []models.Ticket {
	return models.Tickets
}

func CreateTicket(ticket models.Ticket) (models.Ticket, error) {
	ticket.ID = strconv.Itoa(rand.Intn(1000000))
	models.Tickets = append(models.Tickets, ticket)
	return ticket, nil
}

func GetTicketByID(id string) (models.Ticket, error) {
	for _, t := range models.Tickets {
		if t.ID == id {
			return t, nil
		}
	}
	return models.Ticket{}, errors.New("ticket not found")
}

func UpdateTicket(id string, newData models.Ticket) (models.Ticket, error) {
	for i, t := range models.Tickets {
		if t.ID == id {
			newData.ID = id
			models.Tickets[i] = newData
			return newData, nil
		}
	}
	return models.Ticket{}, errors.New("ticket not found")
}

func DeleteTicket(id string) error {
	for i, t := range models.Tickets {
		if t.ID == id {
			models.Tickets = append(models.Tickets[:i], models.Tickets[i+1:]...)
			return nil
		}
	}
	return errors.New("ticket not found")
}
