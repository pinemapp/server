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
		&models.Project{},
		&models.ProjectUser{},
		&models.List{},
		&models.Task{},
	)
}
