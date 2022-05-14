package services

import (
	"github.com/zoshigayan/zocms/db"
	"github.com/zoshigayan/zocms/models"
)

func Migrate() {
	db := db.DbManager()
	db.AutoMigrate(&models.Entry{})
}
