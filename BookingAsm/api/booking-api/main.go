package main

import (
	"BookingAsm/api/booking-api/handlers"
	"BookingAsm/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	//Create grpc client connect
	bookingConn, err := grpc.Dial(":2224", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	//Singleton
	bookingClient := pb.NewBookingServiceClient(bookingConn)
	//Handler for GIN Gonic
	h := handlers.NewBookingHandler(bookingClient)
	g := gin.Default()

	//creat routes
	gr := g.Group("/v3/api")
	gr.POST("/creat", h.CreatBooking)
	gr.POST("/cancel", h.CancelBooking)
	gr.GET("/viewbycode", h.ViewBookingByCode)
	//Listen and serve
	g.Run(":3333")
}