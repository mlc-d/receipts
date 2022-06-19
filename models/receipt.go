package models

import (
	"recibosV2/data"
	"recibosV2/errorh"
	"time"

	"gorm.io/gorm"
)

type Receipt struct {
	Id        uint16         `gorm:"primaryKey" json:"id,omitempty"`
	ServiceID uint8          `json:"service_id,omitempty"`
	Service   Service        `json:"service,omitempty"`
	Amount    uint16         `json:"amount,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func CreateReceipt(r Receipt) {
	db := data.GormDb
	result := db.Create(&r)
	errorh.Handle(result.Error)
}

func GetReceipt(r Receipt) (receipt Receipt) {
	db := data.GormDb
	db.Where("id = ?", r.Id).Find(&receipt)
	return
}

func GetReceipts() (receipts []Receipt) {
	db := data.GormDb
	db.Find(&receipts)
	return
}

func UpdateReceipt(r Receipt) {
	db := data.GormDb
	db.Model(&r).Updates(r)
}

func DeleteReceipt(r Receipt) {
	db := data.GormDb
	db.Delete(&r)
}
