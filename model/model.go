package model

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jinzhu/gorm"
	"fmt"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "/db/sqlite.db")
	if err != nil {
		fmt.Println(err)
		panic("connect postgres failed")
	}
	db.LogMode(true)
	DB = db

	return db
}
