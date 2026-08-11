Feature: Parking vehicles
  As a parking lot operator
  I want vehicles to be parked and unparked according to capacity rules
  So that the lot never overbooks and drivers get a valid ticket

  Scenario: Successfully parking a vehicle
    Given a parking lot with capacity for 1 vehicle
    When I park a car with license plate "ABC-123"
    Then the parking should succeed
    And I should receive a valid ticket

  Scenario: Parking is rejected when the lot is full
    Given a parking lot with capacity for 1 vehicle
    And the lot already has 1 vehicle parked
    When I park a car with license plate "XYZ-999"
    Then the parking should be rejected
    And I should see the error "parking: no available slot for this vehicle type"

  Scenario: Successfully unparking a vehicle
    Given a parking lot with capacity for 1 vehicle
    And the lot already has 1 vehicle parked
    When I unpark the vehicle using its ticket
    Then the unparking should succeed

  Scenario: Unparking an unknown ticket fails
    Given a parking lot with capacity for 1 vehicle
    When I unpark a vehicle using ticket ID "does-not-exist"
    Then the unparking should be rejected
    And I should see the error "parking: vehicle not found in lot"