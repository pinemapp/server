package db

import (
	"github.com/jinzhu/gorm"
	"github.com/pinem/server/models"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Team{},
		&models.TeamUser{},
		&models.Board{},
		&models.BoardUser{},
		&models.List{},
		&models.Task{},
	)
}
