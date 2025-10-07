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

func (s *ReservationService) Reserve(userID, resourceID uint, start time.Time, end time.Time) (*Reservation, error) {
	if start.Compare(end) != -1 || start.Compare(time.Now()) != +1 {
		return nil, errors.New("invalid dates")
	}

	// TODO check valid resource

	if ov, err := s.repo.CountOverlapping(resourceID, start, end); ov > 0 {
		return nil, err
	}

	r := &Reservation{
		//TODO
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

// TODO: func (s *ReservationService) ListUserReservations(userID uint, filters...) ([]Reservation, error) {}
