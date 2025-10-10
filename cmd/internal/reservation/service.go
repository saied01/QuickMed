package reservation

import (
	"errors"
	"time"
)

type ReservationService struct {
	repo *ReservationRepository
}

func NewReservationService(rep *ReservationRepository) *ReservationService {
	return &ReservationService{repo: rep}
}

// Valida start < end, duración mínima/máx., ventana de reserva (no reservar en el pasado), formato/timezones.
//
// Verifica existencia del Resource y su capacidad/condiciones.
//
// Chequea solapamientos (usando repo.FindOverlapping).
//
// Crea reserva dentro de una transacción para evitar condiciones de carrera.
//
// Devuelve ErrConflict si hay solapamiento.

func (s *ReservationService) Create(userID, resourceID uint, start time.Time, end time.Time) (*Reservation, error) {
	if start.Compare(end) != -1 || start.Compare(time.Now()) != +1 {
		return nil, errors.New("invalid dates")
	}

	// TODO check valid resource

	if ov, err := s.repo.CountOverlapping(resourceID, start, end); ov > 0 {
		return nil, err
	}

	tx := s.repo.BeginTx()
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer tx.Rollback()

	ov, err := s.repo.CountOverlappingTx(tx, resourceID, start, end)
	if err != nil {
		return nil, err
	}
	if ov > 0 {
		return nil, errors.New("overlapping conflict")
	}

	r := &Reservation{
		UserID:            userID,
		ResourceID:        resourceID,
		StartTime:         start,
		EndTime:           end,
		CreatedAt:         time.Now(),
		ReservationStatus: "booked",
	}

	if err := s.repo.CreateTx(tx, r); err != nil {
		return nil, err
	}
	if t := tx.Commit(); t.Error != nil {
		return nil, t.Error
	}

	return r, nil
}

// Verifica permisos (propietario o rol admin).
//
// Cambia Status a cancelled. Opcional: notifica al usuario.

// func (s *ReservationService) CancelReservation(userID, reservationID uint) error {}

// Marcar como finished. (usualmente por sistema o job)

// func (s *ReservationService) FinishReservation(reservationID uint) error {}

// Devuelve huecos libres, útil para UI.

// func (s *ReservationService) GetAvailability(resourceID uint, from, to time.Time) ([]FreeSlot, error) {
// }

func (s *ReservationService) GetUserReservations(userID uint) ([]Reservation, error) {
	return s.repo.FindByUserID(userID)
}
