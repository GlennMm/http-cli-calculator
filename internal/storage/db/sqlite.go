package db

import (
	"os"

	"goalculator/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dB *gorm.DB

func initDB() *gorm.DB {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(dir+"/.config/goalculator", 0o755)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open(dir+"/.config/goalculator/history.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Operation{})

	return db
}

func GetDb() *gorm.DB {
	if dB == nil {
		dB = initDB()
	}
	return dB
}
