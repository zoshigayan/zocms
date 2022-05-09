package models

import "time"

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
