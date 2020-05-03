package models

import (
	"github.com/jinzhu/gorm"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("Sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&Book{})
	return db
}
