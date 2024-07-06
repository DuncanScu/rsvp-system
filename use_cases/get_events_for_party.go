package usecases

import (
	"fmt"

	"github.com/DuncanScu/rsvp_system/entities"
	"github.com/DuncanScu/rsvp_system/use_cases/repositories"
)

type GetEventsByParty struct {
	EventsRepo repositories.EventRepository
	PartyRepo  repositories.PartyRepository
}

func (uc *GetEventsByParty) Execute(uniqueId string) ([]*entities.Event, error) {
	party, err := uc.PartyRepo.GetPartyByUniqueId(uniqueId)
	if err != nil {
		return nil, err
	}
	var events []*entities.Event
	for _, event := range party.Events {
		e, err := uc.EventsRepo.GetEventById(event.ID)
		if err != nil {
			return nil, err
		}
		fmt.Println(e.Name)
		events = append(events, e)
	}

	return events, err
}
