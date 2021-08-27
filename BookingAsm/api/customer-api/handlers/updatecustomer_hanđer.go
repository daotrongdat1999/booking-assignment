package handlers

import (
	"BookingAsm/api/customer-api/requestes"
	"BookingAsm/api/customer-api/responses"
	"BookingAsm/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (ch *customerHandler) UpdateCustomer(c *gin.Context) {
	req := requestes.UpdateCustomerRequest{}

	//bind form to req
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

	//proto request
	pRequest := pb.Customer{
		CustomerName: req.Name,
		Address:      req.Address,
		Phone:        req.Phone,
		LicenseId:    req.License_id,
		Id:           req.ID,
		Email:        req.Email,
		Active:       req.Active,
	}

	//call update on grpc
	pResponse, err := ch.customerClient.UpdateCustomer(c, &pRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	//dto contain result to return client
	dto := responses.CustomerResponse{
		ID:         pResponse.Id,
		Name:       pResponse.CustomerName,
		Phone:      pResponse.Phone,
		License_id: pResponse.LicenseId,
		Address:    pResponse.Address,
		Email:      pResponse.Email,
		Active:     pResponse.Active,
	}

	//json to response http
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}