package billing

import (
	"time"

	"github.com/NigusA-12/parking-lot/internal/vehicle"
)

// FeeStrategy defines the contract for calculating a parking fee
// given a duration. Each vehicle type (or pricing rule) implements
// this independently, so new pricing rules can be added without
// modifying existing ones — the Strategy Pattern applying Open/Closed.
type FeeStrategy interface {
	CalculateFee(duration time.Duration) float64
}

// CarFeeStrategy charges a flat hourly rate for cars.
type CarFeeStrategy struct{}

func (CarFeeStrategy) CalculateFee(duration time.Duration) float64 {
	const hourlyRate = 2.0
	return duration.Hours() * hourlyRate
}

// BikeFeeStrategy charges a flat hourly rate for bikes.
type BikeFeeStrategy struct{}

func (BikeFeeStrategy) CalculateFee(duration time.Duration) float64 {
	const hourlyRate = 0.5
	return duration.Hours() * hourlyRate
}

// TruckFeeStrategy charges a flat hourly rate for trucks.
type TruckFeeStrategy struct{}

func (TruckFeeStrategy) CalculateFee(duration time.Duration) float64 {
	const hourlyRate = 5.0
	return duration.Hours() * hourlyRate
}

// strategyFor maps a vehicle type to its corresponding FeeStrategy.
// This is the ONLY place in the system that knows which vehicle type
// uses which pricing rule — everywhere else works through the
// FeeStrategy interface polymorphically.
func strategyFor(vType vehicle.VehicleType) FeeStrategy {
	switch vType {
	case vehicle.Car:
		return CarFeeStrategy{}
	case vehicle.Bike:
		return BikeFeeStrategy{}
	case vehicle.Truck:
		return TruckFeeStrategy{}
	default:
		return CarFeeStrategy{} // sensible default; revisit if new types are added without a strategy
	}
}
