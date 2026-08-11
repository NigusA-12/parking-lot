package parking

import "errors"

// Sentinel errors for the parking domain. Callers can compare against
// these using errors.Is() instead of matching on error message strings.
var (
	ErrSlotNotAvailable = errors.New("parking: no available slot for this vehicle type")
	ErrVehicleNotFound  = errors.New("parking: vehicle not found in lot")
	ErrInvalidVehicle   = errors.New("parking: invalid vehicle")
)
