package main

import (
	"BookingAsm/grpc/customer-grpc/handlers"
	"BookingAsm/grpc/customer-grpc/repositories"
	"BookingAsm/pb"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	bookingConnect, err := grpc.Dial(":2224", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	bookingClient := pb.NewBookingServiceClient(bookingConnect)

	listen, err := net.Listen("tcp", ":2222")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	customerRepos, err := repositories.NewDBManager()
	if err != nil {
		panic(err)
	}

	h, err := handlers.NewCustomerHandler(customerRepos, bookingClient)
	if err != nil {
		panic(err)
	}

	reflection.Register(s)
	pb.RegisterCustomerServiceServer(s, h)

	fmt.Println("Listen at port: 2222")

	s.Serve(listen)
}