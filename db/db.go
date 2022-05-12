package db

import (
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

	log.Println("DB connection succeessed.")
}

func DbManager() *gorm.DB {
	return db
}
