package parking

// Repository defines the storage contract for parking tickets.
// Business logic depends on this interface, never on a concrete
// storage implementation — that is the Dependency Inversion at work.
type Repository interface {
	// Save persists a new ticket.
	Save(ticket *Ticket) error

	// FindByID retrieves a ticket by its ID.
	// Returns ErrVehicleNotFound if no matching ticket exists.
	FindByID(ticketID string) (*Ticket, error)

	// Delete removes a ticket by its ID.
	// Returns ErrVehicleNotFound if no matching ticket exists.
	Delete(ticketID string) error

	// Count returns the number of currently stored (active) tickets.
	Count() (int, error)
}
