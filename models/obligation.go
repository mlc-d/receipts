package models

import (
	"recibosV2/data"
	"recibosV2/errorh"

	"gorm.io/gorm"
)

type Obligation struct {
	Id        uint16         `bun:"primaryKey" json:"id,omitempty"`
	PersonID  uint8          `json:"person_id,omitempty"`
	Person    Person         `json:"person,omitempty"`
	ServiceID uint8          `json:"service_id,omitempty"`
	Service   Service        `json:"service,omitempty"`
	DeletedAt gorm.DeletedAt `bun:"index" json:"deleted_at,omitempty"`
}

func CreateObligation(o Obligation) {
	db := data.Gdb
	result := db.Create(&o)
	errorh.Handle(result.Error)
}

func GetObligations() (obligations []Obligation) {
	db := data.Gdb
	db.Find(&obligations)
	return
}

func UpdateObligation(o Obligation) {
	db := data.Gdb
	db.Model(&o).Updates(&o)
}

func DeleteObligation(o Obligation) {
	db := data.Gdb
	db.Delete(&o)
}
