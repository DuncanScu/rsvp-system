package entities

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name     string
	Time     time.Time
	Location string
	Parties  []*Party `gorm:"many2many:party_events;"`
}
