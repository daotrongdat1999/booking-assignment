package handlers

import (
	"BookingAsm/pb"
	"net/http"

	bReq "BookingAsm/api/booking-api/requestes"
	bRes "BookingAsm/api/booking-api/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (bh *bookingHandler) CancelBooking(c *gin.Context) {
	req := bReq.CancelBookingRequest{}

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

	pRequest := pb.CancelBookingRequest{
		Code:   req.Code,
		Status: req.Status,
	}

	pResponse, err := bh.bookingClient.CancelBooking(c.Request.Context(), &pRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &bRes.BookingResponse{
		ID:          pResponse.Id,
		Customer_id: pResponse.CustomerId,
		Flight_id:   pResponse.FlightId,
		Code:        pResponse.Code,
		Status:      pResponse.Status,
		Booked_date: pResponse.GetBookedDate().AsTime(),
	}

	//return to client
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}
