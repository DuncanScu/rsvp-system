package entities

import (
	"github.com/DuncanScu/rsvp_system/entities/enums/status"
	"gorm.io/gorm"
)

type RSVP struct {
	gorm.Model
	EventId uint `gorm:"uniqueIndex:idx_event_user"`
	UserId  uint `gorm:"uniqueIndex:idx_event_user"`
	Status  status.StatusType
}
