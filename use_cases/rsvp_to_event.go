package usecases

import (
	"github.com/DuncanScu/rsvp_system/entities/enums/status"
	"github.com/DuncanScu/rsvp_system/use_cases/repositories"
)

type RsvpToEvent struct {
	RsvpRepo repositories.RsvpRepo
}

func (uc *RsvpToEvent) Execute(eventId, userId uint, status status.StatusType) error {
	rsvp, err := uc.RsvpRepo.GetRsvpByEventAndUser(eventId, userId)
	if err != nil {
		return err
	}

	rsvp.Status = status
	err = uc.RsvpRepo.UpdateRsvp(rsvp)
	return err
}
