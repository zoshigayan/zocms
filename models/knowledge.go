package models

import (
	"github.com/zoshigayan/zocms/db"
	"time"
)

type Knowledge struct {
	Slug        string `gorm:"PrimaryKey"`
	Title       string
	BodyMD      string
	BodyHTML    string
	Draft       bool `gorm:"default:true"`
	PublishedAt time.Time
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

func KnowledgeAll() []Knowledge {
	cursor := db.DbManager()
	knowledges := []Knowledge{}
	cursor.Find(&knowledges)
	return knowledges
}

func KnowledgeFind(slug string) Knowledge {
	cursor := db.DbManager()
	knowledge := Knowledge{}
	cursor.First(&knowledge, "slug = ?", slug)
	return knowledge
}

func KnowledgeUpdate(slug string, value Knowledge) {
	cursor := db.DbManager()
	knowledge := KnowledgeFind(slug)
	cursor.Model(&knowledge).Updates(value)
}

func KnowledgeCreate(value Knowledge) {
	cursor := db.DbManager()
	cursor.Create(&value)
}
