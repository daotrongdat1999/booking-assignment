package repositories

import (
	"BookingAsm/grpc/customer-grpc/models"
	"context"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error)
	FindCustomer(ctx context.Context, id uuid.UUID) (*models.Customer, error)
	UpdateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error)
}

type dbManager struct {
	*gorm.DB
}

func NewDBManager() (CustomerRepository, error) {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=daotrongdat99 dbname=customer port=5433 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Customer{})

	return &dbManager{db}, nil
}

//creat a customer
func (m *dbManager) CreateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

//find a customer by id
func (m *dbManager) FindCustomer(ctx context.Context, id uuid.UUID) (*models.Customer, error) {
	var result models.Customer
	err := m.Where(&models.Customer{ID: id}).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

//update customer
func (m *dbManager) UpdateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error) {
	err := m.Where(&models.Customer{ID: model.ID}).Updates(model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}
