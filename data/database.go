package data

import (
	"recibosV2/error"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	// SqlDb sql.DB
	GormDb gorm.DB
)

func init() {
	GormDb, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	error.Handle(err)
}
