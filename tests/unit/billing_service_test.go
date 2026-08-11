package unit

import (
	"testing"
	"time"

	"github.com/NigusA-12/parking-lot/internal/billing"
	"github.com/NigusA-12/parking-lot/internal/parking"
	"github.com/NigusA-12/parking-lot/internal/vehicle"
)

func TestBillingService_CalculateFee_Car(t *testing.T) {
	svc := billing.NewBillingService()
	ticket := &parking.Ticket{
		Vehicle:   vehicle.NewCar("ABC-123"),
		EntryTime: time.Now().Add(-2 * time.Hour), // parked 2 hours ago
	}

	fee, err := svc.CalculateFee(ticket)

	if err != nil {
		t.Fatalf("expected no error calculating fee, got: %v", err)
	}
	// Car rate is $2.00/hour, parked ~2 hours -> ~$4.00
	if fee < 3.9 || fee > 4.1 {
		t.Errorf("expected fee close to 4.00 for a 2-hour car park, got %.2f", fee)
	}
}

func TestBillingService_CalculateFee_Bike(t *testing.T) {
	svc := billing.NewBillingService()
	ticket := &parking.Ticket{
		Vehicle:   vehicle.NewBike("BIKE-001"),
		EntryTime: time.Now().Add(-4 * time.Hour), // parked 4 hours ago
	}

	fee, err := svc.CalculateFee(ticket)

	if err != nil {
		t.Fatalf("expected no error calculating fee, got: %v", err)
	}
	// Bike rate is $0.50/hour, parked ~4 hours -> ~$2.00
	if fee < 1.9 || fee > 2.1 {
		t.Errorf("expected fee close to 2.00 for a 4-hour bike park, got %.2f", fee)
	}
}

func TestBillingService_CalculateFee_Truck(t *testing.T) {
	svc := billing.NewBillingService()
	ticket := &parking.Ticket{
		Vehicle:   vehicle.NewTruck("TRUCK-777"),
		EntryTime: time.Now().Add(-1 * time.Hour), // parked 1 hour ago
	}

	fee, err := svc.CalculateFee(ticket)

	if err != nil {
		t.Fatalf("expected no error calculating fee, got: %v", err)
	}
	// Truck rate is $5.00/hour, parked ~1 hour -> ~$5.00
	if fee < 4.9 || fee > 5.1 {
		t.Errorf("expected fee close to 5.00 for a 1-hour truck park, got %.2f", fee)
	}
}

func TestBillingService_CalculateFee_NilTicket(t *testing.T) {
	svc := billing.NewBillingService()

	_, err := svc.CalculateFee(nil)

	if err == nil {
		t.Fatal("expected an error when calculating fee for a nil ticket")
	}
}
