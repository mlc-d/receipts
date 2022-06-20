package models

import (
	"recibosV2/data"
	"recibosV2/errorh"

	"gorm.io/gorm"
)

type FixedFee struct {
	Id        uint16         `bun:"primaryKey" json:"id,omitempty"`
	PersonID  uint8          `json:"person_id,omitempty"`
	Person    Person         `json:"person,omitempty"`
	ServiceID uint8          `json:"service_id,omitempty"`
	Service   Service        `json:"service,omitempty"`
	Amount    float32        `json:"amount,omitempty"`
	DeletedAt gorm.DeletedAt `bun:"index" json:"deleted_at,omitempty"`
}

func CreateFixedFee(f FixedFee) {
	db := data.Gdb
	result := db.Create(&f)
	errorh.Handle(result.Error)
}

func GetFixedFees() (fixedFee []FixedFee) {
	db := data.Gdb
	db.Find(&fixedFee)
	return
}

func UpdateFixedFee(f FixedFee) {
	db := data.Gdb
	db.Model(&f).Updates(&f)
}

func DeleteFixedFee(f FixedFee) {
	db := data.Gdb
	db.Delete(&f)
}
