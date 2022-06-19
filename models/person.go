package models

import (
	"recibosV2/data"
	"recibosV2/errorh"

	"gorm.io/gorm"
)

type Person struct {
	Id        uint8          `gorm:"primaryKey" json:"id,omitempty"`
	Name      string         `gorm:"unique,type:varchar(30)" json:"name,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func CreatePerson(p Person) {
	db := data.GormDb
	result := db.Create(&p)
	errorh.Handle(result.Error)
}

func GetPerson(p Person) (person Person) {
	db := data.GormDb
	db.Where("name = ?", p.Name).Find(&person)
	return
}

func GetPersons() (persons []Person) {
	db := data.GormDb
	db.Find(&persons)
	return
}

func UpdatePerson(p Person) {
	db := data.GormDb
	db.Model(&p).Updates(p)
}

func DeletePerson(p Person) {
	db := data.GormDb
	db.Delete(&p)
}
