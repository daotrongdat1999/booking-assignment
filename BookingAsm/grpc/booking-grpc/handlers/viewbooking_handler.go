package handlers

import (
	"BookingAsm/pb"
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *BookingHandler) SearchBooking(ctx context.Context, in *pb.SearchBookingRequest) (*pb.BookingInfor, error) {
	booking, err := h.bookingRepository.FindBooking(ctx, in.Code)
	if err != nil {
		return nil, err
	}

	//call customer gRPC to search customer infor
	customer, err := h.customerClient.FindCustomer(ctx, &pb.FindCustomerRequest{
		Id: booking.Customer_id.String(),
	})
	if err != nil {
		return nil, err
	}

	//call flight gRPC to search flight infor
	flight, err := h.flightClient.SearchFlight(ctx, &pb.SearchFlightRequest{
		Id: booking.Flight_id.String(),
	})
	if err != nil {
			return nil, err
	}

	//return booking infor contain customer infor and flight infor
	return &pb.BookingInfor{
		BookingDetail:  &pb.Booking{
			Id:         booking.ID.String(),
			CustomerId: booking.Customer_id.String(),
			FlightId:   booking.Flight_id.String(),
			Code:       booking.Code,
			Status:     booking.Status,
			BookedDate: timestamppb.New(booking.Booked_date),
		},
		FlightDetail:   flight,
		CustomerDetail: customer,
	}, nil
}
