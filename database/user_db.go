package database

import (
	"log/slog"

	"github.com/DuncanScu/rsvp_system/entities"
	"gorm.io/gorm"
)

type UserDb struct {
	Db *gorm.DB
}

func (u *UserDb) GetUsers() ([]entities.User, error) {
	slog.Info("Retrieving users from db")

	var user entities.User
	result := u.Db.First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	users := []entities.User{user}
	return users, nil
}

func (u *UserDb) GetUser(id int) (entities.User, error) {
	slog.Info("Retrieving user from db")
	var user entities.User
	result := u.Db.First(&user, "user_id = ?", id)
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return user, nil
}

func (u *UserDb) AddUser(user entities.User) (entities.User, error) {
	u.Db.Create(&user)
	return user, nil
}
