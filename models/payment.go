package models

import (
	"recibosV2/data"
	"recibosV2/errorh"
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	Id        uint32         `gorm:"primaryKey" json:"id,omitempty"`
	ReceiptID uint16         `json:"receipt_id,omitempty"`
	Receipt   Receipt        `json:"receipt,omitempty"`
	PersonID  uint8          `json:"person_id,omitempty"`
	Person    Person         `json:"person,omitempty"`
	Amount    float32        `gorm:"index" json:"amount,omitempty"`
	CreatedAt time.Time      `json:"date"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func CreatePayment(p Payment) {
	db := data.Gdb
	result := db.Create(&p)
	errorh.Handle(result.Error)
}

func GetPayments() (payments []Payment) {
	db := data.Gdb
	db.Find(&payments)
	return
}

func UpdatePayment(p Payment) {
	db := data.Gdb
	db.Model(&p).Updates(&p)
}

func DeletePayment(p Payment) {
	db := data.Gdb
	db.Delete(&p)
}
