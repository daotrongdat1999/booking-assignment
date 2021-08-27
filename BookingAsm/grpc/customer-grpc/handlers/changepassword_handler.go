package handlers

import (
	"BookingAsm/pb"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//Change Password
func (h *CustomerHandler) ChangePassword(ctx context.Context, in *pb.ChangePasswordRequest) (*pb.Customer, error) {
	//found customer will be change password by idzzzz
	customerUpdated, err := h.customerRepository.FindCustomer(ctx, uuid.MustParse(in.Id))
	if err != nil {
		if err == sql.ErrNoRows { //return err not found rows in db
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	//up date Password for customer
	customerUpdated.Password = in.Password

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
		Active:       newCustomer.Active,
		Id:           newCustomer.ID.String(),
	}, nil
}
