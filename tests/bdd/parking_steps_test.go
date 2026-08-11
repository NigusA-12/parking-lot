package bdd

import (
	"fmt"
	"testing"

	"github.com/cucumber/godog"

	"github.com/NigusA-12/parking-lot/internal/repository"
	"github.com/NigusA-12/parking-lot/internal/service"
	"github.com/NigusA-12/parking-lot/internal/vehicle"
)

// parkingTestContext holds state shared across the steps of a single
// scenario. Godog gives each scenario a fresh instance of this via
// InitializeScenario below, so scenarios never leak state into each other.
type parkingTestContext struct {
	svc          *service.ParkingService
	lastTicketID string
	lastErr      error
}

func (tc *parkingTestContext) aParkingLotWithCapacityFor(capacity int) error {
	repo := repository.NewInMemoryParkingRepository()
	tc.svc = service.NewParkingService(repo, capacity)
	return nil
}

func (tc *parkingTestContext) theLotAlreadyHasVehicleParked(count int) error {
	for i := 0; i < count; i++ {
		plate := fmt.Sprintf("PRE-FILL-%d", i)
		ticket, err := tc.svc.Park(vehicle.NewCar(plate))
		if err != nil {
			return fmt.Errorf("setup failed while pre-filling the lot: %w", err)
		}
		tc.lastTicketID = ticket.ID
	}
	return nil
}

func (tc *parkingTestContext) iParkACarWithLicensePlate(plate string) error {
	ticket, err := tc.svc.Park(vehicle.NewCar(plate))
	tc.lastErr = err
	if ticket != nil {
		tc.lastTicketID = ticket.ID
	}
	return nil
}

func (tc *parkingTestContext) iUnparkTheVehicleUsingItsTicket() error {
	_, tc.lastErr = tc.svc.Unpark(tc.lastTicketID)
	return nil
}

func (tc *parkingTestContext) iUnparkAVehicleUsingTicketID(ticketID string) error {
	_, tc.lastErr = tc.svc.Unpark(ticketID)
	return nil
}

func (tc *parkingTestContext) theParkingShouldSucceed() error {
	if tc.lastErr != nil {
		return fmt.Errorf("expected parking to succeed, got error: %w", tc.lastErr)
	}
	return nil
}

func (tc *parkingTestContext) iShouldReceiveAValidTicket() error {
	if tc.lastTicketID == "" {
		return fmt.Errorf("expected a valid ticket ID, got empty string")
	}
	return nil
}

func (tc *parkingTestContext) theParkingShouldBeRejected() error {
	if tc.lastErr == nil {
		return fmt.Errorf("expected parking to be rejected, but it succeeded")
	}
	return nil
}

func (tc *parkingTestContext) theUnparkingShouldSucceed() error {
	if tc.lastErr != nil {
		return fmt.Errorf("expected unparking to succeed, got error: %w", tc.lastErr)
	}
	return nil
}

func (tc *parkingTestContext) theUnparkingShouldBeRejected() error {
	if tc.lastErr == nil {
		return fmt.Errorf("expected unparking to be rejected, but it succeeded")
	}
	return nil
}

func (tc *parkingTestContext) iShouldSeeTheError(expectedMessage string) error {
	if tc.lastErr == nil {
		return fmt.Errorf("expected an error with message %q, got no error", expectedMessage)
	}
	if tc.lastErr.Error() != expectedMessage {
		return fmt.Errorf("expected error message %q, got %q", expectedMessage, tc.lastErr.Error())
	}
	return nil
}

// InitializeScenario registers every step definition and ensures each
// scenario gets its own fresh parkingTestContext, so state never leaks
// between scenarios.
func InitializeScenario(sc *godog.ScenarioContext) {
	tc := &parkingTestContext{}

	sc.Given(`^a parking lot with capacity for (\d+) vehicle$`, tc.aParkingLotWithCapacityFor)
	sc.Given(`^the lot already has (\d+) vehicle parked$`, tc.theLotAlreadyHasVehicleParked)
	sc.When(`^I park a car with license plate "([^"]*)"$`, tc.iParkACarWithLicensePlate)
	sc.When(`^I unpark the vehicle using its ticket$`, tc.iUnparkTheVehicleUsingItsTicket)
	sc.When(`^I unpark a vehicle using ticket ID "([^"]*)"$`, tc.iUnparkAVehicleUsingTicketID)
	sc.Then(`^the parking should succeed$`, tc.theParkingShouldSucceed)
	sc.Then(`^I should receive a valid ticket$`, tc.iShouldReceiveAValidTicket)
	sc.Then(`^the parking should be rejected$`, tc.theParkingShouldBeRejected)
	sc.Then(`^the unparking should succeed$`, tc.theUnparkingShouldSucceed)
	sc.Then(`^the unparking should be rejected$`, tc.theUnparkingShouldBeRejected)
	sc.Then(`^I should see the error "([^"]*)"$`, tc.iShouldSeeTheError)
}

// TestFeatures is the entry point go test uses to run our Gherkin
// scenarios through Godog.
func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"../../features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
