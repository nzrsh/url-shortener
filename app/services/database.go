package services

import (
	_ "github.com/gofiber/fiber/v3/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var SqlDB *gorm.DB

func CreateOrConnectDB(path string) error {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return err
	}
	db = db.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Info)})

	err = db.AutoMigrate(&Link{})
	if err != nil {
		return err
	}
	SqlDB = db
	return nil
}
