package initializers

import (
	"time"

	"github.com/DuncanScu/rsvp_system/entities"
	"gorm.io/gorm"
)

func SeedDb() {
	duncan := &entities.User{
		Model:       gorm.Model{ID: 1},
		FirstName:   "Duncan",
		LastName:    "Scuffell",
		Email:       "",
		PhoneNumber: "",
	}

	stacey := &entities.User{
		Model:       gorm.Model{ID: 2},
		FirstName:   "Stacey",
		LastName:    "Scuffell",
		Email:       "",
		PhoneNumber: "",
	}

	DB.Create(duncan)
	DB.Create(stacey)

	scuffellParty := &entities.Party{
		UniqueID: "VVoveui9QC",
		Name:     "Scuffell Party",
		PartyMembers: []*entities.User{
			duncan,
			stacey,
		}}

	DB.Create(scuffellParty)

	receptionEvent := &entities.Event{
		Name:     "Reception",
		Time:     time.Now().Add(3),
		Location: "The Venue",
		Parties: []*entities.Party{
			scuffellParty,
		},
	}

	DB.Create(receptionEvent)

	rsvpMe := &entities.RSVP{
		UserId: 1,
		Status: "Pending",
	}

	rsvpStacey := &entities.RSVP{
		UserId: 2,
		Status: "Pending",
	}

	DB.Create(rsvpMe)
	DB.Create(rsvpStacey)
}
