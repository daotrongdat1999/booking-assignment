package main

import (
	"BookingAsm/api/flight-api/handlers"
	"BookingAsm/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	//Create grpc client connect
	customerConn, err := grpc.Dial(":2223", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	//Singleton
	peopleClient := pb.NewFlightServiceClient(customerConn)

	//Handler for GIN Gonic
	h := handlers.NewFlightHandler(peopleClient)
	g := gin.Default()

	//Create routes
	gr := g.Group("/v2/api")
	gr.POST("/create", h.CreatFlight)
	gr.POST("/update", h.UpdateFlight)
	gr.GET("/find", h.SearchFlight)
	//Listen and serve
	g.Run(":3333")
}