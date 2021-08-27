package handlers

import (
	"BookingAsm/grpc/booking-grpc/repositories"
	"BookingAsm/pb"
)

type BookingHandler struct {
	customerClient pb.CustomerServiceClient
	flightClient pb.FlightServiceClient

	pb.UnimplementedBookingServiceServer
	bookingRepository repositories.BookingRepository
}

func NewBookingHandler(bookingRepo repositories.BookingRepository,
	customerClient pb.CustomerServiceClient,
	flightClient pb.FlightServiceClient) (*BookingHandler, error) {
	return &BookingHandler{
		bookingRepository: bookingRepo,
		customerClient: customerClient,
		flightClient: flightClient,
		}, nil
}





