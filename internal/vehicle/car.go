package vehicle

// CarVehicle represents a car parked in the lot.
type CarVehicle struct {
	plate string
}

// NewCar constructs a CarVehicle. Construction is explicit and validated
// here rather than allowing a bare struct literal with an empty plate.
func NewCar(plate string) *CarVehicle {
	return &CarVehicle{plate: plate}
}

func (c *CarVehicle) LicensePlate() string {
	return c.plate
}

func (c *CarVehicle) Type() VehicleType {
	return Car
}
