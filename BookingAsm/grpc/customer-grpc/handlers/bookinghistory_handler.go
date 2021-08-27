package handlers

import (
	"BookingAsm/pb"
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//find history booking by id customer
func (h *CustomerHandler) BookingHistory(ctx context.Context, in *pb.BookingHistoryRequest) (*pb.BookingHistoryResponse, error) {
	bookings, err := h.bookingClient.SearchBookingId(ctx, &pb.SearchBookingByIdRequest{
		Id: in.Id,
	})
	if err != nil {
		if err == sql.ErrNoRows { //return err not found rows in db
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	var pResponse pb.BookingHistoryResponse
	//
	pResponse.CustomerId = in.Id
	for _, v := range bookings.BookingList{
		pResponse.Code = append(pResponse.Code, v.Code)
		pResponse.Id = append(pResponse.Id, v.Id)
		pResponse.FlightId = append(pResponse.FlightId, v.FlightId)
		pResponse.Status = append(pResponse.Status, v.Status)
		pResponse.BookedDate = append(pResponse.BookedDate, v.BookedDate)
	}

	
	return &pResponse, nil
}