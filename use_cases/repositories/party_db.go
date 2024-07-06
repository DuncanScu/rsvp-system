package repositories

import "github.com/DuncanScu/rsvp_system/entities"

type PartyRepository interface {
	GetPartyMembers(partyId string) ([]*entities.User, error)
	GetPartyByUniqueId(uniqueId string) (*entities.Party, error)
}
