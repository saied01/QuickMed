// Package reservation defines the domain model for reservations
// made by users, including their status, schedule, and references
// to reserved resources.
package reservation

import (
	"time"
)

type ReservationStatus string

const (
	StatusBooked    ReservationStatus = "booked"
	StatusCancelled ReservationStatus = "cancelled"
	StatusFinished  ReservationStatus = "finished"
)

type Resource struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"size:80;not null"`
	Type         string `gorm:"size:80;not null"`
	Capacity     uint
	MetadataJSON string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Reservation struct {
	ID                uint `gorm:"primaryKey"`
	ResourceID        uint
	Resource          Resource `gorm:"foreignKey:ResourceID"`
	UserID            uint
	User              uint `gorm:"foreignKey:UserID"`
	StartTime         time.Time
	EndTime           time.Time
	ReservationStatus ReservationStatus `gorm:"type:varchar(20);not null"`
	CreatedAt         time.Time
}
