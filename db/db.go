package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pinem/server/config"
)

var ORM *gorm.DB

func InitORM() error {
	conf := config.Get()
	db, err := gorm.Open("postgres", conf.DbString())
	if err != nil {
		return err
	}
	if err := db.DB().Ping(); err != nil {
		return err
	}
	if conf.ENV != "production" {
		db.LogMode(true)
	}
	ORM = db
	return nil
}
