// Package user defines the domain model and entities related to
// application users. It includes personal data, credentials, and
// the relationship between a user and their reservations.
package user

import (
	"database/sql"
	"time"

	"quickmed/internal/reservation"
)

type User struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:80;not null"`
	Email       string `gorm:"uniqueIndex;size:100;not null"`
	Age         uint8  `gorm:"not null"`
	Password    string `gorm:"size:255;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ActivatedAt sql.NullTime

	Reservations []reservation.Reservation `gorm:"foreignKey:UserID"`

	Specialty     *string
	LicenseNumber *string
	IsVerified    *bool
}
