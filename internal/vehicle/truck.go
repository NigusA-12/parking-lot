package vehicle

// TruckVehicle represents a truck parked in the lot.
type TruckVehicle struct {
	plate string
}

func NewTruck(plate string) *TruckVehicle {
	return &TruckVehicle{plate: plate}
}

func (t *TruckVehicle) LicensePlate() string {
	return t.plate
}

func (t *TruckVehicle) Type() VehicleType {
	return Truck
}
