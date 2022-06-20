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
	Amount    float32        `json:"amount,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func CreateReceipt(r Receipt) {
	db := data.Gdb
	result := db.Create(&r)
	errorh.Handle(result.Error)
}

func GetReceipt(r Receipt) (receipt Receipt) {
	db := data.Gdb
	db.Where("id = ?", r.Id).Find(&receipt)
	return
}

func GetReceipts() (receipts []Receipt) {
	db := data.Gdb
	db.Find(&receipts)
	return
}

func UpdateReceipt(r Receipt) {
	db := data.Gdb
	db.Model(&r).Updates(r)
}

func DeleteReceipt(r Receipt) {
	db := data.Gdb
	db.Delete(&r)
}
