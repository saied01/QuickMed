package reservation

import (
	"gorm.io/gorm"
)

type ReservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository(database *gorm.DB) *ReservationRepository {
	return &ReservationRepository{db: database}
}

func (r *ReservationRepository) Create(res *Reservation) error {
	result := r.db.Create(res)
	return result.Error
}

func (r *ReservationRepository) GetByID(id uint) (*Reservation, error) {
	var reservation Reservation
	result := r.db.First(reservation, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &reservation, nil
}

// ListByUser(userID uint, limit, offset int) ([]Reservation, error)

// ListByResource(resourceID uint, from, to time.Time) ([]Reservation, error)

// FindOverlapping(resourceID uint, start, end time.Time) ([]Reservation, error)

// Update(res *Reservation) error

// Delete(res *Reservation) error

// CountOverlapping(resourceID uint, start, end time.Time) (int64, error) (útil para check rápido)

// Optional: FindAvailableResources(from, to time.Time, capacity uint) ([]Resource, error)
