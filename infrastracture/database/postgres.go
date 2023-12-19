package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func NewPostgresConn() (*gorm.DB, error) {
	return Connect()
}

func Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DatabaseHost"), os.Getenv("DatabasePort"), os.Getenv("DatabaseUser"), os.Getenv("DatabasePassword"), os.Getenv("DatabaseName"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: %v")
	}

	return db, err
}
