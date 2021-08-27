package handlers

import (
	"BookingAsm/pb"

	"github.com/gin-gonic/gin"
)

type BookingHandler interface {
	CreatBooking(c *gin.Context)
	CancelBooking(c *gin.Context)
	ViewBookingByCode(c *gin.Context)
	// ViewAllBooking(c *gin.Context)
}

type bookingHandler struct {
	bookingClient pb.BookingServiceClient
}

func NewBookingHandler(bookingClient pb.BookingServiceClient) BookingHandler {
	return &bookingHandler{
		bookingClient: bookingClient,
	}
}

// //view all booking by id cusomer
// func (bh *bookingHandler) ViewAllBooking(c *gin.Context){
// 	req := requestes.ViewAllBookingById{}

// 	//parse form request
// 	if err := c.ShouldBindJSON(&req); err != nil{
// 		//validate form
// 		if validateErr, ok := err.(validator.ValidationErrors); ok {
// 			errMessage := make([]string, 0)
// 			for _, v := range validateErr {
// 				errMessage = append(errMessage, v.Error())
// 			}

// 			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 				"status": http.StatusText(http.StatusBadRequest),
// 				"error":  errMessage,
// 			})

// 			return
// 	}}

// 	pRequest := &pb.SearchBookingByIdRequest{
// 		Id: req.Customer_id,
// 	}

// 	pResponse, err := bh.bookingClient.SearchBookingId(c.Request.Context(), pRequest)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 			"status": http.StatusText(http.StatusInternalServerError),
// 			"error":  err.Error(),
// 		})
// 		return
// 	}

// 	//mảng booking trả về
// 	dto := make([]responses.BookingResponse,0)

// 	for _, v := range pResponse.BookingList{
// 		tmp := &responses.BookingResponse{
// 			ID:          v.Id,
// 			Customer_id: v.CustomerId,
// 			Flight_id:   v.FlightId,
// 			Code:        v.Code,
// 			Status:      v.Status,
// 			Booked_date: v.BookedDate.AsTime(),
// 		}
// 		dto =append(dto, *tmp)
// 	}

// 	//return to client
// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  http.StatusText(http.StatusOK),
// 		"payload": dto,
// 	})
// }
