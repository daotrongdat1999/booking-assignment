package handlers

import (
	"BookingAsm/grpc/customer-grpc/repositories"
	"BookingAsm/pb"
)

type CustomerHandler struct {
	bookingClient pb.BookingServiceClient
	pb.UnimplementedCustomerServiceServer
	customerRepository repositories.CustomerRepository
}

func NewCustomerHandler(customerRepo repositories.CustomerRepository,
	bookingClient pb.BookingServiceClient) (*CustomerHandler, error) {
	return &CustomerHandler{
		customerRepository: customerRepo,
		bookingClient:      bookingClient,
	}, nil
}







