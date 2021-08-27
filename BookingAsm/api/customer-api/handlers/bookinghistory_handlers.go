package handlers

import (
	"BookingAsm/api/customer-api/requestes"
	"BookingAsm/api/customer-api/responses"
	"BookingAsm/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (ch *customerHandler) BookingHistory(c *gin.Context){
	req := requestes.FindCustomerRequest{}

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

	pCRequest := &pb.BookingHistoryRequest{
		Id: req.ID,
	}

	pResponse, err := ch.customerClient.BookingHistory(c.Request.Context(), pCRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := responses.HistoryResponse{}
	dto.Customrer_id = pResponse.CustomerId

	for _, v := range pResponse.BookedDate{
		dto.Booking_date = append(dto.Booking_date, v.AsTime())
	}
	dto.Booking_code = append(dto.Booking_code, pResponse.Code...)
	dto.Flight_id = append(dto.Flight_id, pResponse.FlightId...)
	dto.Booking_status = append(dto.Booking_status, pResponse.Status...)
	dto.Booking_id = append(dto.Booking_id, pResponse.Id...)
	

	//json to response http
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}