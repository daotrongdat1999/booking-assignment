package handlers

import (
	"BookingAsm/pb"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//update a flight
func (h *FlightHandler) UpdateFlight(ctx context.Context, in *pb.Flight) (*pb.Flight, error) {
	//find the flight will be updated
	flightUpdate, err := h.flightRepository.SearchFlight(ctx, uuid.MustParse(in.Id))
	if err != nil {
		if err == sql.ErrNoRows { //return err not found rows in db
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	//up date fields for flight has just been found
	if in.Name != "" {
		flightUpdate.Name = in.Name
	}

	if in.From != "" {
		flightUpdate.From = in.From
	}

	if in.To != "" {
		flightUpdate.To = in.To
	}

	if in.Date != nil {
		flightUpdate.Date = in.GetDate().AsTime()
	}

	if in.Status != "" {
		flightUpdate.Status = in.Status
	}

	if in.AvailableSlot != int32(flightUpdate.Available_slot) {
		flightUpdate.Available_slot = int(in.AvailableSlot)
	}

	newFlight, err := h.flightRepository.UpdateFlight(ctx, flightUpdate)
	if err != nil {
		return nil, err
	}

	return &pb.Flight{
		Id:            newFlight.ID.String(),
		Name:          newFlight.Name,
		From:          newFlight.From,
		To:            newFlight.To,
		Date:          timestamppb.New(newFlight.Date),
		Status:        newFlight.Status,
		AvailableSlot: int32(newFlight.Available_slot),
	}, nil
}
