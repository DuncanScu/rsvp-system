package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	PartyID     uint
}
