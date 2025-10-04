package reservation

import (
	"time"

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
	result := r.db.First(&reservation, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &reservation, nil
}

func (r *ReservationRepository) ListByUser(userID uint, limit, offset int) ([]Reservation, error) {
	var res []Reservation

	result := r.db.Where("user_id = ?", userID).
		Limit(limit).Offset(offset).Find(&res)
	return res, result.Error
}

func (r *ReservationRepository) ListByResource(resourceID uint, from, to time.Time) ([]Reservation, error) {
	var res []Reservation

	result := r.db.Where("user_id = ? AND start_time >= ? AND end_time <= ?", resourceID, from, to).Find(res)

	return res, result.Error
}

func (r *ReservationRepository) Update(res *Reservation) error {
	return r.db.Save(res).Error
}

func (r *ReservationRepository) Delete(res *Reservation) error {
	result := r.db.Delete(res)
	return result.Error
}

func (r *ReservationRepository) CountOverlapping(resourceID uint, start, end time.Time) (int64, error) {
	var count int64

	result := r.db.Model(&Reservation{}).
		Where("resource_id =? AND NOT (end_time <= ? OR start_time >= ?)", resourceID, start, end).
		Count(&count)

	return count, result.Error
}

// Optional: FindAvailableResources(from, to time.Time, capacity uint) ([]Resource, error)
