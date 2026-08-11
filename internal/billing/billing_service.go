package billing

import (
	"errors"
	"time"

	"github.com/NigusA-12/parking-lot/internal/parking"
)

// ErrInvalidTicket is returned when fee calculation is attempted
// on a nil or otherwise invalid ticket.
var ErrInvalidTicket = errors.New("billing: invalid ticket")

// BillingService calculates fees for parking tickets, delegating the
// actual rate calculation to the appropriate FeeStrategy.
type BillingService struct{}

// NewBillingService constructs a BillingService.
func NewBillingService() *BillingService {
	return &BillingService{}
}

// CalculateFee returns the fee owed for the given ticket, based on
// how long the vehicle has been parked and its vehicle type.
func (s *BillingService) CalculateFee(ticket *parking.Ticket) (float64, error) {
	if ticket == nil || ticket.Vehicle == nil {
		return 0, ErrInvalidTicket
	}

	duration := time.Since(ticket.EntryTime)
	strategy := strategyFor(ticket.Vehicle.Type())

	return strategy.CalculateFee(duration), nil
}
