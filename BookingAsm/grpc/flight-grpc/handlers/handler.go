package handlers

import (
	"BookingAsm/grpc/flight-grpc/repositories"
	"BookingAsm/pb"
)

type FlightHandler struct {
	pb.UnimplementedFlightServiceServer
	flightRepository repositories.FlightRepository
}

func NewFlightHandler(flightRepo repositories.FlightRepository) (*FlightHandler, error) {
	return &FlightHandler{flightRepository: flightRepo}, nil
}
