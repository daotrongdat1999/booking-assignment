package handlers

import (
	"BookingAsm/grpc/flight-grpc/models"
	"BookingAsm/pb"
	"context"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//create a flight
func (h *FlightHandler) CreatFlight(ctx context.Context, in *pb.Flight) (*pb.Flight, error) {
	flight, err := h.flightRepository.CreatFlight(ctx, &models.Flight{
		ID:             uuid.New(),
		Name:           in.Name,
		From:           in.From,
		To:             in.To,
		Date:           in.GetDate().AsTime(),
		Status:         in.Status,
		Available_slot: int(in.AvailableSlot),
	})
	if err != nil {
		return nil, err
	}

	//return flight has just been created
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
