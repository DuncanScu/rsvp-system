package initializers

import (
	"log/slog"

	"github.com/DuncanScu/rsvp_system/entities"
)

func MigrateDb() {
	slog.Info("Migrating db")
	DB.AutoMigrate(
		&entities.RSVP{},
		&entities.Event{},
		&entities.Party{},
		&entities.User{},
	)
	slog.Info("Db migration complete")
}
