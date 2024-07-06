package repositories

import "github.com/DuncanScu/rsvp_system/entities"

type EventRepository interface {
	GetEventById(eventId uint) (*entities.Event, error)
}
