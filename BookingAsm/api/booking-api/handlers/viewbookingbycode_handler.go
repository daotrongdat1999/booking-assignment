package handlers

import (
	bRes "BookingAsm/api/booking-api/responses"
	cRes "BookingAsm/api/customer-api/responses"
	fRes "BookingAsm/api/flight-api/responses"

	bReq "BookingAsm/api/booking-api/requestes"

	"BookingAsm/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (bh *bookingHandler) ViewBookingByCode(c *gin.Context) {
	req := &bReq.SearchBookingRequest{}

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

	pBRequest := &pb.SearchBookingRequest{
		Code: req.Code,
	}

	pBResponse, err := bh.bookingClient.SearchBooking(c.Request.Context(), pBRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &bRes.SearchBookingResponse{
		BookingResponse: bRes.BookingResponse{
			ID:          pBResponse.BookingDetail.Id,
			Customer_id: pBResponse.BookingDetail.CustomerId,
			Flight_id:   pBResponse.FlightDetail.Id,
			Code:        pBResponse.BookingDetail.Code,
			Status:      pBResponse.BookingDetail.Status,
			Booked_date: pBResponse.BookingDetail.BookedDate.AsTime(),
		},
		CustomerInfor: cRes.CustomerResponse{
			ID:         pBResponse.BookingDetail.CustomerId,
			Name:       pBResponse.CustomerDetail.CustomerName,
			Phone:      pBResponse.CustomerDetail.Phone,
			License_id: pBResponse.CustomerDetail.LicenseId,
			Address:    pBResponse.CustomerDetail.Address,
			Email:      pBResponse.CustomerDetail.Email,
			Active:     pBResponse.CustomerDetail.Active,
		},
		FlightInfor: fRes.FlightResponse{
			ID:            pBResponse.BookingDetail.FlightId,
			Name:          pBResponse.FlightDetail.Name,
			From:          pBResponse.FlightDetail.From,
			To:            pBResponse.FlightDetail.To,
			Date:          pBResponse.FlightDetail.Date.AsTime(),
			Status:        pBResponse.FlightDetail.Status,
			AvailableSlot: pBResponse.FlightDetail.AvailableSlot,
		},
	}

	//return to client
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}
