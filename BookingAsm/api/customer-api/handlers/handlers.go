package handlers

import (
	"BookingAsm/pb"

	"github.com/gin-gonic/gin"
)

type CustomerHandler interface {
	CreatCustomer(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	ChangePassword(c *gin.Context)
	FindCustomer(c *gin.Context)
	BookingHistory(c *gin.Context)
}

type customerHandler struct {
	customerClient pb.CustomerServiceClient
}

func NewCustomerHandler(customerClient pb.CustomerServiceClient) CustomerHandler {
	return &customerHandler{
		customerClient: customerClient,
	}
}








