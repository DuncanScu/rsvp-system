package database

import (
	"github.com/DuncanScu/rsvp_system/entities"
	"gorm.io/gorm"
)

type EventDb struct {
	Db *gorm.DB
}

func (e *EventDb) GetEventById(eventId uint) (*entities.Event, error) {
	var event *entities.Event
	err := e.Db.First(&event, &entities.Event{Model: gorm.Model{ID: eventId}}).Error
	if err != nil {
		return nil, err
	}
	return event, nil
}
