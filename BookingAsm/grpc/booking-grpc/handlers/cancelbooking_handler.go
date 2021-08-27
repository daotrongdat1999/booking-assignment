package handlers

import (
	"BookingAsm/pb"
	"context"
	
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *BookingHandler) CancelBooking(ctx context.Context, in *pb.CancelBookingRequest) (*pb.Booking, error) {
	//find the booking to cancel
	bookingCancel, err := h.bookingRepository.FindBooking(ctx, in.Code)
	if err != nil {
		return nil, err
	}

	//change status to cancel
	if in.Status != "" {
		bookingCancel.Status = in.Status
	}

	newBooking, err := h.bookingRepository.CancelBooking(ctx, bookingCancel)
	if err != nil {
		return nil, err
	}

	//return booking has just been cancel
	return &pb.Booking{
		Id:         newBooking.ID.String(),
		CustomerId: newBooking.Customer_id.String(),
		FlightId:   newBooking.Flight_id.String(),
		Code:       newBooking.Code,
		Status:     newBooking.Status,
		BookedDate: timestamppb.New(newBooking.Booked_date),
	}, nil
}