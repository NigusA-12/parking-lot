package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/NigusA-12/parking-lot/internal/billing"
	"github.com/NigusA-12/parking-lot/internal/handler"
	"github.com/NigusA-12/parking-lot/internal/repository"
	"github.com/NigusA-12/parking-lot/internal/service"
)

const defaultLotCapacity = 20

func main() {
	repo := repository.NewInMemoryParkingRepository()
	parkingService := service.NewParkingService(repo, defaultLotCapacity)
	billingService := billing.NewBillingService()
	parkingHandler := handler.NewParkingHandler(parkingService, billingService)

	router := gin.Default()
	parkingHandler.RegisterRoutes(router)

	log.Println("Parking Lot System listening on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
