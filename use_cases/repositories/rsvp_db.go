package repositories

import (
	"github.com/DuncanScu/rsvp_system/entities"
)

type RsvpRepo interface {
	GetRsvpByEventAndUser(eventId, userId uint) (*entities.RSVP, error)
	UpdateRsvp(rsvp *entities.RSVP) error
}
