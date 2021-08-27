package main

import (
	"BookingAsm/api/customer-api/handlers"
	"BookingAsm/pb"
	// "net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	//Create grpc client connect
	customerConn, err := grpc.Dial(":2222", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	//Singleton
	peopleClient := pb.NewCustomerServiceClient(customerConn)

	//Handler for GIN Gonic
	h := handlers.NewCustomerHandler(peopleClient)
	g := gin.Default()

	//Create routes
	gr := g.Group("/v1/api")
	gr.POST("/create", h.CreatCustomer)
	gr.POST("/update", h.UpdateCustomer)
	gr.POST("/changepass", h.ChangePassword)
	gr.GET("/find", h.FindCustomer)
	//Listen and serve
	g.Run(":3333")
}