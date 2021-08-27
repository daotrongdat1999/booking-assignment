package handlers

import (
	"BookingAsm/pb"

	"github.com/gin-gonic/gin"
)

type FlightHandler interface {
	CreatFlight(c *gin.Context)
	SearchFlight(c *gin.Context)
	UpdateFlight(c *gin.Context)
}

type flightHandler struct {
	flightClient pb.FlightServiceClient
}

func NewFlightHandler(flightClient pb.FlightServiceClient) FlightHandler {
	return &flightHandler{
		flightClient: flightClient,
	}
}
