package handlers

import (
	"BookingAsm/pb"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//update customer
func (h *CustomerHandler) UpdateCustomer(ctx context.Context, in *pb.Customer) (*pb.Customer, error) {
	//found customer will be updated by id
	customerUpdated, err := h.customerRepository.FindCustomer(ctx, uuid.MustParse(in.Id))
	if err != nil {
		if err == sql.ErrNoRows { //return err not found rows in db
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	//up date fields for customer has just been found
	if in.CustomerName != "" {
		customerUpdated.Name = in.CustomerName
	}

	if in.Address != "" {
		customerUpdated.Address = in.Address
	}

	if in.Phone != "" {
		customerUpdated.Phone = in.Phone
	}

	if in.LicenseId != "" {
		customerUpdated.License_id = in.LicenseId
	}

	if in.Email != "" {
		customerUpdated.Email = in.Email
	}

	if in.Active != customerUpdated.Active {
		customerUpdated.Active = in.Active
	}

	//update to db
	newCustomer, err := h.customerRepository.UpdateCustomer(ctx, customerUpdated)
	if err != nil {
		return nil, err
	}

	return &pb.Customer{
		CustomerName: newCustomer.Name,
		Address:      newCustomer.Address,
		Phone:        newCustomer.Phone,
		LicenseId:    newCustomer.License_id,
		Email:        newCustomer.Email,
		Active:       newCustomer.Active,
		Id:           newCustomer.ID.String(),
	}, nil
}
