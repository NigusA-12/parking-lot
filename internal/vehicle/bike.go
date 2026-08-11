package vehicle

// BikeVehicle represents a bike parked in the lot.
type BikeVehicle struct {
	plate string
}

func NewBike(plate string) *BikeVehicle {
	return &BikeVehicle{plate: plate}
}

func (b *BikeVehicle) LicensePlate() string {
	return b.plate
}

func (b *BikeVehicle) Type() VehicleType {
	return Bike
}
