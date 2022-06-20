package models

import (
	"recibosV2/data"
	"recibosV2/errorh"

	"gorm.io/gorm"
)

type Service struct {
	Id        uint8          `gorm:"primaryKey" json:"id,omitempty"`
	Name      string         `gorm:"unique,type:varchar(30)" json:"name,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func CreateService(s Service) {
	db := data.Gdb
	result := db.Create(&s)
	errorh.Handle(result.Error)
}

func GetService(s Service) (service Service) {
	db := data.Gdb
	db.Where("name = ?", s.Name).Find(&service)
	return
}

func GetServices() (services []Service) {
	db := data.Gdb
	db.Find(&services)
	return
}

func UpdateService(s Service) {
	db := data.Gdb
	db.Model(&s).Updates(s)
}

func DeleteService(s Service) {
	db := data.Gdb
	db.Delete(&s)
}
