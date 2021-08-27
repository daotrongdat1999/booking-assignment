package handlers

import (
	"BookingAsm/grpc/booking-grpc/models"
	"BookingAsm/pb"
	"context"

	"github.com/google/uuid"
	"github.com/rs/xid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *BookingHandler) CreatBooking(ctx context.Context, in *pb.Booking) (*pb.Booking, error) {
	booking, err := h.bookingRepository.CreatBooking(ctx, &models.Booking{
		ID:          uuid.New(),
		Customer_id: uuid.MustParse(in.CustomerId),
		Flight_id:   uuid.MustParse(in.FlightId),
		Code:        xid.New().String(),
		Status:      in.Status,
		Booked_date: in.GetBookedDate().AsTime(),
	})
	if err != nil {
		return nil, err
	}

	//return flight has just been created
	return &pb.Booking{
		Id:         booking.ID.String(),
		CustomerId: booking.Customer_id.String(),
		FlightId:   booking.Flight_id.String(),
		Code:       booking.Code,
		Status:     booking.Status,
		BookedDate: timestamppb.New(booking.Booked_date),
	}, nil
}