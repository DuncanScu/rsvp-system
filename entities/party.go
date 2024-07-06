package entities

import "gorm.io/gorm"

type Party struct {
	gorm.Model
	UniqueID     string `gorm:"unique"`
	Name         string
	PartyMembers []*User
	Events       []*Event `gorm:"many2many:party_events;"`
}
