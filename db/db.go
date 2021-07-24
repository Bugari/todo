package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func OpenDb(locator string) {
	var err error
	Conn, err = gorm.Open(sqlite.Open(locator), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Conn.AutoMigrate(&Task{})
}
