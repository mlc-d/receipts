package data

import (
	"recibosV2/errorh"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	// SqlDb sql.DB
	Gdb *gorm.DB
)

func Connect() {
	var err error
	Gdb, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	errorh.Handle(err)
}
