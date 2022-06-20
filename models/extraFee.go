package models

import (
	"recibosV2/data"
	"recibosV2/errorh"

	"gorm.io/gorm"
)

type ExtraFee struct {
	Id        uint16         `bun:"primaryKey" json:"id,omitempty"`
	PersonID  uint8          `json:"person_id,omitempty"`
	Person    Person         `json:"person,omitempty"`
	ServiceID uint8          `json:"service_id,omitempty"`
	Service   Service        `json:"service,omitempty"`
	Amount    float32        `json:"amount,omitempty"`
	DeletedAt gorm.DeletedAt `bun:"index" json:"deleted_at,omitempty"`
}

func CreateExtraFee(f ExtraFee) {
	db := data.Gdb
	result := db.Create(&f)
	errorh.Handle(result.Error)
}

func GetExtraFees() (ExtraFee []ExtraFee) {
	db := data.Gdb
	db.Find(&ExtraFee)
	return
}

func UpdateExtraFee(f ExtraFee) {
	db := data.Gdb
	db.Model(&f).Updates(&f)
}

func DeleteExtraFee(f ExtraFee) {
	db := data.Gdb
	db.Delete(&f)
}
