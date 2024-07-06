package database

import (
	"github.com/DuncanScu/rsvp_system/entities"
	"gorm.io/gorm"
)

type RsvpDb struct {
	Db *gorm.DB
}

func (db *RsvpDb) GetRsvpByEventAndUser(eventId, userId uint) (*entities.RSVP, error) {
	var rsvp *entities.RSVP
	err := db.Db.First(&rsvp, &entities.RSVP{EventId: eventId, UserId: userId}).Error

	if err != nil {
		return nil, err
	}

	return rsvp, nil
}

func (db *RsvpDb) UpdateRsvp(rsvp *entities.RSVP) error {
	err := db.Db.Model(rsvp).Update("Status", rsvp.Status).Error
	return err
}
