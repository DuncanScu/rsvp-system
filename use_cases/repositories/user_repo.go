package repositories

import "github.com/DuncanScu/rsvp_system/entities"

type UserRepository interface {
	GetUsers() ([]entities.User, error)
	AddUser(entities.User) (entities.User, error)
	GetUser(int) (entities.User, error)
}
