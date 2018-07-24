package db

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pinem/server/config"
)

var ORM *gorm.DB

func InitORM() error {
	conf := config.Get()
	gorm.NowFunc = func() time.Time {
		return time.Now().In(conf.GetLocation())
	}

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
	migrate(db)

	return nil
}

func Transaction(db *gorm.DB, cb func(*gorm.DB) error) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := cb(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
