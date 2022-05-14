package models

import (
	"github.com/zoshigayan/zocms/db"
	"time"
)

type Entry struct {
	ID          uint `gorm:"PrimaryKey;autoIncrement"`
	Slug        string
	Type        string `gorm:"index"`
	Title       string
	BodyMD      string
	BodyHTML    string
	Draft       bool `gorm:"default:true"`
	PublishedAt time.Time
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

func EntryAll() []Entry {
	cursor := db.DbManager()
	entries := []Entry{}
	cursor.Find(&entries)
	return entries
}

func EntryFind(id uint) Entry {
	cursor := db.DbManager()
	entry := Entry{}
	cursor.First(&entry, id)
	return entry
}

func EntryUpdate(id uint, value Entry) Entry {
	cursor := db.DbManager()
	entry := EntryFind(id)
	cursor.Model(&entry).Updates(value)
	return entry
}

func EntryCreate(value Entry) Entry {
	cursor := db.DbManager()
	cursor.Create(&value)
	return value
}
