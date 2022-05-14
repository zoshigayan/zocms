package models

import (
	"github.com/zoshigayan/zocms/db"
)

func KnowledgeAll() []Entry {
	cursor := db.DbManager()
	knowledges := []Entry{}
	cursor.Where("type = 'knowledge'").Find(&knowledges)
	return knowledges
}

func KnowledgeFind(id uint) Entry {
	cursor := db.DbManager()
	knowledge := Entry{}
	cursor.First(&knowledge, id)
	return knowledge
}

func KnowledgeUpdate(id uint, value Entry) {
	cursor := db.DbManager()
	knowledge := KnowledgeFind(id)
	cursor.Model(&knowledge).Updates(value)
}

func KnowledgeCreate(value Entry) {
	cursor := db.DbManager()
	value.Type = "knowledge"
	cursor.Create(&value)
}
