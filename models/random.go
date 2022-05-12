package models

import (
	"github.com/zoshigayan/zocms/db"
	"time"
)

type Random struct {
	ID          uint `gorm:"PrimaryKey;autoIncrement"`
	Title       string
	BodyMD      string
	BodyHTML    string
	Draft       bool `gorm:"default:true"`
	PublishedAt time.Time
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

func RandomAll() []Random {
	cursor := db.DbManager()
	randoms := []Random{}
	cursor.Find(&randoms)
	return randoms
}

func RandomFind(id uint) Random {
	cursor := db.DbManager()
	random := Random{}
	cursor.First(&random, id)
	return random
}

func RandomUpdate(id uint, value Random) Random {
	cursor := db.DbManager()
	random := RandomFind(id)
	cursor.Model(&random).Updates(value)
	return random
}

func RandomCreate(value Random) Random {
	cursor := db.DbManager()
	cursor.Create(&value)
	return value
}
