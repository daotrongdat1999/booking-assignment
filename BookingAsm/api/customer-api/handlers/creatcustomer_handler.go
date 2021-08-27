package handlers

import (
	"BookingAsm/api/customer-api/requestes"
	"BookingAsm/api/customer-api/responses"
	"BookingAsm/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (ch *customerHandler) CreatCustomer(c *gin.Context) {
	req := requestes.CreatCustomerRequest{} //declare CreatCustomerRequest object

	//parse form from http request
	if err := c.ShouldBindJSON(&req); err != nil {
		//validate form
		if validateErr, ok := err.(validator.ValidationErrors); ok {
			errMessage := make([]string, 0)
			for _, v := range validateErr {
				errMessage = append(errMessage, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessage,
			})

			return
		}
	}

	//protobuf model customer
	pRequest := &pb.Customer{
		CustomerName: req.Name,
		Address:      req.Address,
		Phone:        req.Phone,
		LicenseId:    req.License_id,
		Active:       true,
		Email:        req.Email,
		Password:     req.Password,
	}

	//gRPC client call creat new customer
	pResponse, err := ch.customerClient.CreateCustomer(c.Request.Context(), pRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	//response object with field same pResponse
	dto := &responses.CustomerResponse{
		Name:       pResponse.CustomerName,
		Phone:      pResponse.Phone,
		License_id: pResponse.LicenseId,
		Address:    pResponse.Address,
		Email:      pResponse.Email,
		ID:         pResponse.Id,
	}

	//json to response http
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}
