package db

import (
	"github.com/zoshigayan/zocms/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB
var err error

func Init() {
	db, err = gorm.Open(sqlite.Open("zocms.db"), &gorm.Config{})
	if err != nil {
		panic("DB connection failed.")
	}

	db.AutoMigrate(&models.Knowledge{})
	// db.AutoMigrate(&models.Random{})
	log.Println("DB connection succeessed.")
}

func DbManager() *gorm.DB {
	return db
}
