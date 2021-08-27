package handlers

import (
	"BookingAsm/pb"
	"context"

	"github.com/google/uuid"
)

//find customer by id
func (h *CustomerHandler) FindCustomer(ctx context.Context, in *pb.FindCustomerRequest) (*pb.Customer, error) {
	customer, err := h.customerRepository.FindCustomer(ctx, uuid.MustParse(in.Id))
	if err != nil {
		return nil, err
	}

	//return customer has just been found
	return &pb.Customer{
		CustomerName: customer.Name,
		Address:      customer.Address,
		Phone:        customer.Phone,
		LicenseId:    customer.License_id,
		Active:       customer.Active,
		Id:           customer.ID.String(),
		Email:        customer.Email,
		Password:     customer.Password,
	}, nil
}