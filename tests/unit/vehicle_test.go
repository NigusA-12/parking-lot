package unit

import (
	"testing"

	"github.com/NigusA-12/parking-lot/internal/vehicle"
)

func TestNewVehicle_Car(t *testing.T) {
	v, err := vehicle.NewVehicle(vehicle.Car, "ABC-123")

	if err != nil {
		t.Fatalf("expected no error creating a car, got: %v", err)
	}
	if v.Type() != vehicle.Car {
		t.Errorf("expected type %q, got %q", vehicle.Car, v.Type())
	}
	if v.LicensePlate() != "ABC-123" {
		t.Errorf("expected plate %q, got %q", "ABC-123", v.LicensePlate())
	}
}

func TestNewVehicle_Bike(t *testing.T) {
	v, err := vehicle.NewVehicle(vehicle.Bike, "BIKE-01")

	if err != nil {
		t.Fatalf("expected no error creating a bike, got: %v", err)
	}
	if v.Type() != vehicle.Bike {
		t.Errorf("expected type %q, got %q", vehicle.Bike, v.Type())
	}
}

func TestNewVehicle_Truck(t *testing.T) {
	v, err := vehicle.NewVehicle(vehicle.Truck, "TRK-01")

	if err != nil {
		t.Fatalf("expected no error creating a truck, got: %v", err)
	}
	if v.Type() != vehicle.Truck {
		t.Errorf("expected type %q, got %q", vehicle.Truck, v.Type())
	}
}

func TestNewVehicle_UnknownType(t *testing.T) {
	v, err := vehicle.NewVehicle(vehicle.VehicleType("SPACESHIP"), "XYZ-999")

	if err == nil {
		t.Fatal("expected an error for an unknown vehicle type, got none")
	}
	if v != nil {
		t.Fatalf("expected nil vehicle on error, got: %+v", v)
	}
}
