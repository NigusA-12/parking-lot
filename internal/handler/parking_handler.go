package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/NigusA-12/parking-lot/internal/billing"
	"github.com/NigusA-12/parking-lot/internal/parking"
	"github.com/NigusA-12/parking-lot/internal/service"
	"github.com/NigusA-12/parking-lot/internal/vehicle"
)

// ParkingHandler translates HTTP requests into ParkingService/BillingService
// calls, and their results back into HTTP responses. It contains no
// business logic of its own.
type ParkingHandler struct {
	service *service.ParkingService
	billing *billing.BillingService
}

// NewParkingHandler constructs a ParkingHandler with injected
// ParkingService and BillingService dependencies, following the same
// Dependency Injection pattern used throughout this project.
func NewParkingHandler(svc *service.ParkingService, billingSvc *billing.BillingService) *ParkingHandler {
	return &ParkingHandler{service: svc, billing: billingSvc}
}

// RegisterRoutes wires this handler's endpoints onto the given
// Gin router. Keeping route registration here (rather than in main)
// keeps all of this handler's HTTP surface in one place.
func (h *ParkingHandler) RegisterRoutes(router *gin.Engine) {
	router.POST("/park", h.Park)
	router.POST("/unpark/:ticketId", h.Unpark)
}

// parkRequest is the expected JSON body for POST /park.
type parkRequest struct {
	LicensePlate string `json:"licensePlate" binding:"required"`
	VehicleType  string `json:"vehicleType" binding:"required"`
}

// parkResponse is returned on successful parking.
type parkResponse struct {
	TicketID     string `json:"ticketId"`
	LicensePlate string `json:"licensePlate"`
}

// unparkResponse is returned on successful unparking, including the
// fee owed for the completed parking session.
type unparkResponse struct {
	Message string  `json:"message"`
	Fee     float64 `json:"fee"`
}

// Park handles POST /park.
func (h *ParkingHandler) Park(c *gin.Context) {
	var req parkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	v, err := vehicle.NewVehicle(vehicle.VehicleType(req.VehicleType), req.LicensePlate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticket, err := h.service.Park(v)
	if err != nil {
		writeServiceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, parkResponse{
		TicketID:     ticket.ID,
		LicensePlate: ticket.Vehicle.LicensePlate(),
	})
}

// Unpark handles POST /unpark/:ticketId.
func (h *ParkingHandler) Unpark(c *gin.Context) {
	ticketID := c.Param("ticketId")

	ticket, err := h.service.Unpark(ticketID)
	if err != nil {
		writeServiceError(c, err)
		return
	}

	fee, err := h.billing.CalculateFee(ticket)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to calculate fee"})
		return
	}

	c.JSON(http.StatusOK, unparkResponse{
		Message: "vehicle unparked successfully",
		Fee:     fee,
	})
}

// writeServiceError maps a domain sentinel error to the appropriate
// HTTP status code and writes the response.
func writeServiceError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, parking.ErrInvalidVehicle):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case errors.Is(err, parking.ErrSlotNotAvailable):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	case errors.Is(err, parking.ErrVehicleNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}
