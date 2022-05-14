package models

import (
	"github.com/zoshigayan/zocms/db"
)

func RandomAll() []Entry {
	cursor := db.DbManager()
	randoms := []Entry{}
	cursor.Where("type = 'random'").Find(&randoms)
	return randoms
}

func RandomFind(id uint) Entry {
	cursor := db.DbManager()
	random := Entry{}
	cursor.First(&random, id)
	return random
}

func RandomUpdate(id uint, value Entry) Entry {
	cursor := db.DbManager()
	random := RandomFind(id)
	cursor.Model(&random).Updates(value)
	return random
}

func RandomCreate(value Entry) Entry {
	cursor := db.DbManager()
	value.Type = "random"
	cursor.Create(&value)
	return value
}
