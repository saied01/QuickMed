// Package db provides database connection utilities using GORM and PostgreSQL.
// It exposes a helper function New that initializes and returns a *gorm.DB
// instance configured with the application’s default connection parameters.
package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func New() *gorm.DB {
	dsn := "host=localhost user=devuser password=devpass dbname=devdb port=5432 sslmode=disable TimeZone=America/Argentina/Buenos_Aires"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to initialize database, got error %v", err)
	}

	return db
}

func Truncate(db *gorm.DB, tableName string) {
	if err := db.Exec("TRUNCATE TABLE " + tableName + " RESTART IDENTITY CASCADE").Error; err != nil {
		log.Fatalf("failed to truncate table %s: %v", tableName, err)
	}
}
