// Package db provides database connection utilities using GORM and PostgreSQL.
// It exposes a helper function New that initializes and returns a *gorm.DB
// instance configured with the application’s default connection parameters.
package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Argentina/Buenos-Aires"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("error connecting to database.")
	}

	return db
}
