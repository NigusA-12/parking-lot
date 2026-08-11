package parking

import (
	"time"

	"github.com/NigusA-12/parking-lot/internal/vehicle"
)

// Ticket represents a single parking event: a vehicle occupying a slot
// from a point in time until it exits.
type Ticket struct {
	ID        string
	Vehicle   vehicle.Vehicle
	SlotID    string
	EntryTime time.Time
}
