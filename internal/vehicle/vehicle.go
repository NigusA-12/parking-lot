package vehicle

import "errors"

// VehicleType identifies the category of a vehicle.
type VehicleType string

const (
	Car   VehicleType = "CAR"
	Bike  VehicleType = "BIKE"
	Truck VehicleType = "TRUCK"
)

// Vehicle is the contract every vehicle type must satisfy.
type Vehicle interface {
	LicensePlate() string
	Type() VehicleType
}

// ErrUnknownVehicleType is returned when NewVehicle is asked to
// construct a vehicle type it doesn't recognize.
var ErrUnknownVehicleType = errors.New("vehicle: unknown vehicle type")

// NewVehicle is the single point of construction for any Vehicle,
// given its type and license plate. Callers (such as the HTTP handler)
// depend only on this factory and the Vehicle interface — never on
// concrete constructors like NewCar/NewBike/NewTruck directly.
func NewVehicle(vType VehicleType, plate string) (Vehicle, error) {
	switch vType {
	case Car:
		return NewCar(plate), nil
	case Bike:
		return NewBike(plate), nil
	case Truck:
		return NewTruck(plate), nil
	default:
		return nil, ErrUnknownVehicleType
	}
}
