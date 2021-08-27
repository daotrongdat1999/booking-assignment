package handlers

import (
	"BookingAsm/api/flight-api/requestes"
	"BookingAsm/api/flight-api/responses"
	"BookingAsm/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (fh *flightHandler) UpdateFlight(c *gin.Context) {
	req := requestes.UpdateFlightRequest{}

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

	pRequest := pb.Flight{
		Id:            req.ID,
		Name:          req.Name,
		From:          req.From,
		To:            req.To,
		Date:          timestamppb.New(req.Date),
		Status:        req.Status,
		AvailableSlot: req.AvailableSlot,
	}

	pResponse, err := fh.flightClient.UpdateFlight(c.Request.Context(), &pRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	//response object with field same pResponse
	dto := &responses.FlightResponse{
		ID:            pResponse.Id,
		Name:          pResponse.Name,
		From:          pResponse.From,
		To:            pResponse.To,
		Date:          pResponse.GetDate().AsTime(),
		Status:        pResponse.Status,
		AvailableSlot: pResponse.AvailableSlot,
	}

	//json to response http
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}
