package unit

import (
	"errors"
	"testing"

	"github.com/NigusA-12/parking-lot/internal/parking"
	"github.com/NigusA-12/parking-lot/internal/repository"
	"github.com/NigusA-12/parking-lot/internal/service"
	"github.com/NigusA-12/parking-lot/internal/vehicle"
)

// newTestService wires up a ParkingService with a fresh in-memory
// repository. Centralizing this setup avoids repeating the same
// three lines in every test function.
func newTestService(capacity int) *service.ParkingService {
	repo := repository.NewInMemoryParkingRepository()
	return service.NewParkingService(repo, capacity)
}

func TestParkingService_Park_Success(t *testing.T) {
	svc := newTestService(1)
	car := vehicle.NewCar("ABC-123")

	ticket, err := svc.Park(car)

	if err != nil {
		t.Fatalf("expected no error parking into an available lot, got: %v", err)
	}
	if ticket == nil {
		t.Fatal("expected a non-nil ticket after successful parking")
	}
	if ticket.Vehicle.LicensePlate() != "ABC-123" {
		t.Errorf("expected ticket vehicle plate %q, got %q", "ABC-123", ticket.Vehicle.LicensePlate())
	}
}

func TestParkingService_Park_Full(t *testing.T) {
	svc := newTestService(1)
	firstCar := vehicle.NewCar("AAA-111")
	secondCar := vehicle.NewCar("BBB-222")

	_, err := svc.Park(firstCar)
	if err != nil {
		t.Fatalf("expected first park to succeed, got error: %v", err)
	}

	_, err = svc.Park(secondCar)
	if !errors.Is(err, parking.ErrSlotNotAvailable) {
		t.Fatalf("expected ErrSlotNotAvailable when lot is full, got: %v", err)
	}
}

func TestParkingService_Park_NilVehicle(t *testing.T) {
	svc := newTestService(1)

	ticket, err := svc.Park(nil)

	if !errors.Is(err, parking.ErrInvalidVehicle) {
		t.Fatalf("expected ErrInvalidVehicle for nil vehicle, got: %v", err)
	}
	if ticket != nil {
		t.Fatalf("expected nil ticket when parking fails, got: %+v", ticket)
	}
}

func TestParkingService_Unpark_Success(t *testing.T) {
	svc := newTestService(1)
	car := vehicle.NewCar("ABC-123")

	ticket, err := svc.Park(car)
	if err != nil {
		t.Fatalf("expected park to succeed, got error: %v", err)
	}

	_, err = svc.Unpark(ticket.ID)
	if err != nil {
		t.Fatalf("expected unpark to succeed, got error: %v", err)
	}

	// The slot should now be free again.
	secondCar := vehicle.NewCar("XYZ-999")
	_, err = svc.Park(secondCar)
	if err != nil {
		t.Fatalf("expected slot to be free after unpark, got error: %v", err)
	}
}

func TestParkingService_Unpark_NotFound(t *testing.T) {
	svc := newTestService(1)

	_, err := svc.Unpark("non-existent-ticket-id")
	if !errors.Is(err, parking.ErrVehicleNotFound) {
		t.Fatalf("expected ErrVehicleNotFound for unknown ticket ID, got: %v", err)
	}
}
