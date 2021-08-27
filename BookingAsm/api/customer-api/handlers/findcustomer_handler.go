package handlers

import (
	"BookingAsm/api/customer-api/requestes"
	"BookingAsm/api/customer-api/responses"
	"BookingAsm/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (ch *customerHandler) FindCustomer(c *gin.Context) {
	req := requestes.FindCustomerRequest{} //declare a FindCustomerRequest object
	
	//parse form request
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

	//protobuf model find customer request
	pRequest := pb.FindCustomerRequest{
		Id: req.ID,
	}

	//gRPC client call find customer
	pResponse, err := ch.customerClient.FindCustomer(c.Request.Context(), &pRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	//response object with field same pResponse
	dto := responses.CustomerResponse{
		ID:         pResponse.Id,
		Name:       pResponse.CustomerName,
		Phone:      pResponse.Phone,
		License_id: pResponse.LicenseId,
		Address:    pResponse.Address,
		Email:      pResponse.Email,
	}

	//json to response client
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}