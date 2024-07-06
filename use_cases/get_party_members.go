package usecases

import (
	"github.com/DuncanScu/rsvp_system/entities"
	"github.com/DuncanScu/rsvp_system/use_cases/repositories"
)

type GetPartyMembers struct {
	PartyRepo repositories.PartyRepository
}

func (uc GetPartyMembers) Execute(partyId string) ([]*entities.User, error) {
	// This isnt really correct. I should get the party object, preload the party members and return them
	return uc.PartyRepo.GetPartyMembers(partyId)
}
