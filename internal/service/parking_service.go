package service

import (
	"time"

	"github.com/NigusA-12/parking-lot/internal/parking"
	"github.com/NigusA-12/parking-lot/internal/vehicle"
)

// ParkingService implements parking business rules. It depends only
// on the parking.Repository interface — never on a concrete storage
// type — which is Dependency Inversion applied in practice.
type ParkingService struct {
	repo     parking.Repository
	capacity int
}

// NewParkingService constructs a ParkingService. The repository is
// injected by the caller (Dependency Injection) rather than created
// internally, so ParkingService never decides how tickets are stored.
func NewParkingService(repo parking.Repository, capacity int) *ParkingService {
	return &ParkingService{
		repo:     repo,
		capacity: capacity,
	}
}

// Park assigns the given vehicle a ticket, occupying one slot.
// It returns ErrInvalidVehicle if v is nil, or ErrSlotNotAvailable
// if the lot has no remaining capacity.
func (s *ParkingService) Park(v vehicle.Vehicle) (*parking.Ticket, error) {
	if v == nil {
		return nil, parking.ErrInvalidVehicle
	}

	count, err := s.repo.Count()
	if err != nil {
		return nil, err
	}
	if count >= s.capacity {
		return nil, parking.ErrSlotNotAvailable
	}

	ticket := &parking.Ticket{
		ID:        generateTicketID(),
		Vehicle:   v,
		EntryTime: time.Now(),
	}

	if err := s.repo.Save(ticket); err != nil {
		return nil, err
	}
	return ticket, nil
}

// Unpark releases the slot held by the ticket with the given ID and
// returns the ticket that was removed, so callers (such as billing)
// can use its details after the slot has been freed.
// It returns ErrVehicleNotFound if no matching ticket exists.
func (s *ParkingService) Unpark(ticketID string) (*parking.Ticket, error) {
	ticket, err := s.repo.FindByID(ticketID)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Delete(ticketID); err != nil {
		return nil, err
	}
	return ticket, nil
}

// generateTicketID produces a simple, sufficiently-unique ID for a
// single-process, in-memory parking lot. Not suitable for distributed
// or persistent scenarios — revisit if/when persistence is introduced.
func generateTicketID() string {
	return time.Now().Format("20060102150405.000000000")
}
