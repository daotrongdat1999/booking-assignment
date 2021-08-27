package handlers

import (
	"BookingAsm/pb"
	"context"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//find flight by id
func (h *FlightHandler) SearchFlight(ctx context.Context, in *pb.SearchFlightRequest) (*pb.Flight, error) {
	flight, err := h.flightRepository.SearchFlight(ctx, uuid.MustParse(in.Id))
	if err != nil {
		return nil, err
	}

	return &pb.Flight{
		Id:            flight.ID.String(),
		Name:          flight.Name,
		From:          flight.From,
		To:            flight.To,
		Date:          timestamppb.New(flight.Date),
		Status:        flight.Status,
		AvailableSlot: int32(flight.Available_slot),
	}, nil
}
