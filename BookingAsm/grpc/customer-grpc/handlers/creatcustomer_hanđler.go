package handlers

import (
	"BookingAsm/pb"
	"context"
	"BookingAsm/grpc/customer-grpc/models"

	"github.com/google/uuid"
)
//creat customer
func (h *CustomerHandler) CreateCustomer(ctx context.Context, in *pb.Customer) (*pb.Customer, error) {
	//creat a new customer to db
	customer, err := h.customerRepository.CreateCustomer(ctx, &models.Customer{
		ID:         uuid.New(),
		Name:       in.CustomerName,
		Address:    in.Address,
		Phone:      in.Phone,
		License_id: in.LicenseId,
		Email:      in.Email,
		Password:   in.Password,
		Active:     in.Active,
	})
	if err != nil {
		return nil, err
	}

	//return customer has just been created
	return &pb.Customer{
		Id:           customer.ID.String(),
		CustomerName: customer.Name,
		Address:      customer.Address,
		Phone:        customer.Phone,
		LicenseId:    customer.License_id,
		Email:        customer.Email,
		Password:     customer.Password,
		Active:       customer.Active,
	}, nil
}