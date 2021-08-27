package handlers

import (

	"BookingAsm/pb"
	"context"
	
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//find all booking of one customer
func (h *BookingHandler) SearchBookingId(ctx context.Context, in *pb.SearchBookingByIdRequest)(*pb.ListBooking, error){
	booking, err := h.bookingRepository.FindBookingById(ctx, uuid.MustParse(in.Id))
	if err != nil {
		return nil, err
	}

	var pResponse pb.ListBooking

	//slice Booking search by ID to return
	for _, i := range *booking{
		pResponse.BookingList = append(pResponse.BookingList, &pb.Booking{
			Id:         i.ID.String(),
			CustomerId: i.Customer_id.String(),
			FlightId:   i.Flight_id.String(),
			Code:       i.Code,
			Status:     i.Status,
			BookedDate: timestamppb.New(i.Booked_date),
		})
	}

	return &pResponse, nil
}