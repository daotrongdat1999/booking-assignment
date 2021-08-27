package main

import (
	"BookingAsm/grpc/flight-grpc/handlers"
	"BookingAsm/grpc/flight-grpc/repositories"
	"BookingAsm/pb"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listen, err := net.Listen("tcp", ":2223")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	flightRepos, err := repositories.NewDBManager()
	if err != nil{
		panic(err)
	}

	h, err := handlers.NewFlightHandler(flightRepos)

	reflection.Register(s)
	pb.RegisterFlightServiceServer(s, h)
	
	fmt.Println("Listen at port: 2223")

	s.Serve(listen)
}