package database

import (
	"github.com/DuncanScu/rsvp_system/entities"
	"gorm.io/gorm"
)

type PartyDb struct {
	Db *gorm.DB
}

func (p *PartyDb) GetPartyMembers(partyId string) ([]*entities.User, error) {
	var party entities.Party
	result := p.Db.Preload("PartyMembers").Where(&entities.Party{UniqueID: partyId}).Find(&party)

	if result.Error != nil {
		return nil, result.Error
	}

	return party.PartyMembers, nil
}

func (p *PartyDb) GetPartyByUniqueId(uniqueId string) (*entities.Party, error) {
	var party entities.Party
	result := p.Db.Preload("Events").Where(&entities.Party{UniqueID: uniqueId}).Find(&party)

	if result.Error != nil {
		return nil, result.Error
	}

	return &party, nil
}
