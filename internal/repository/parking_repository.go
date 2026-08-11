package repository

import (
	"github.com/NigusA-12/parking-lot/internal/parking"
)

// InMemoryParkingRepository is an in-memory implementation of
// parking.Repository. It satisfies the interface entirely through
// a private slice — callers never see this detail, only the contract.
type InMemoryParkingRepository struct {
	tickets []*parking.Ticket
}

var _ parking.Repository = (*InMemoryParkingRepository)(nil)

// NewInMemoryParkingRepository constructs an empty in-memory repository.
func NewInMemoryParkingRepository() *InMemoryParkingRepository {
	return &InMemoryParkingRepository{
		tickets: make([]*parking.Ticket, 0),
	}
}

// Save persists a new ticket.
func (r *InMemoryParkingRepository) Save(ticket *parking.Ticket) error {
	r.tickets = append(r.tickets, ticket)
	return nil
}

// FindByID retrieves a ticket by its ID.
func (r *InMemoryParkingRepository) FindByID(ticketID string) (*parking.Ticket, error) {
	for _, t := range r.tickets {
		if t.ID == ticketID {
			return t, nil
		}
	}
	return nil, parking.ErrVehicleNotFound
}

// Delete removes a ticket by its ID.
func (r *InMemoryParkingRepository) Delete(ticketID string) error {
	for i, t := range r.tickets {
		if t.ID == ticketID {
			r.tickets = append(r.tickets[:i], r.tickets[i+1:]...)
			return nil
		}
	}
	return parking.ErrVehicleNotFound
}

// Count returns the number of currently stored tickets.
func (r *InMemoryParkingRepository) Count() (int, error) {
	return len(r.tickets), nil
}
