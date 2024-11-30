package db

import (
	"gorm.io/gorm"
	"log"
)

type Config struct {
	DSN string
}

var db *gorm.DB

func InitDB(config Config) {
	var err error
	dsn := config.DSN
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("Error while closing DB: %v", err)
			return
		}
		err = sqlDB.Close()
		if err != nil {
			return
		}
		log.Println("Database connection closed.")
	}
}
