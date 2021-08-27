package main

import (
	"BookingAsm/grpc/booking-grpc/handlers"
	"BookingAsm/grpc/booking-grpc/repositories"
	"BookingAsm/pb"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	customerConnect, err := grpc.Dial(":2222", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	flightConnect, err := grpc.Dial(":2223", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	customerClient := pb.NewCustomerServiceClient(customerConnect)
	flightClient := pb.NewFlightServiceClient(flightConnect)

	listen, err := net.Listen("tcp", ":2224")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	bookingRepos, err := repositories.NewDBManager()
	if err != nil {
		panic(err)
	}

	h, err := handlers.NewBookingHandler(bookingRepos, customerClient, flightClient)

	reflection.Register(s)
	pb.RegisterBookingServiceServer(s, h)

	fmt.Println("Listen at port: 2224")

	s.Serve(listen)
}