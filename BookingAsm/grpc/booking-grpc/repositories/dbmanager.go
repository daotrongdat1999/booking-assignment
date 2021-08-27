package repositories

import (
	"BookingAsm/grpc/booking-grpc/models"
	"context"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type BookingRepository interface {
	CreatBooking(ctx context.Context, model *models.Booking) (*models.Booking, error)
	CancelBooking(ctx context.Context, model *models.Booking) (*models.Booking, error)
	FindBooking(ctx context.Context, code string) (*models.Booking, error)
	FindBookingById(ctx context.Context, id uuid.UUID) (*[]models.Booking, error)

}

type dbManager struct{
	*gorm.DB
}

func NewDBManager() (BookingRepository, error) {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=daotrongdat99 dbname=booking port=5433 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Booking{})

	return &dbManager{db}, nil
} 

//creat new booking
func (m *dbManager) CreatBooking(ctx context.Context, model *models.Booking) (*models.Booking, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

//find booking by code
func (m *dbManager) FindBooking(ctx context.Context, code string) (*models.Booking, error){
	var result models.Booking
	err := m.Where(&models.Booking{Code: code}).First(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}

//cancel booking
func (m *dbManager) CancelBooking(ctx context.Context, model *models.Booking) (*models.Booking, error) {
	err := m.Where(&models.Booking{Code: model.Code}).Updates(model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

//find all booking by id customer
func (m *dbManager)FindBookingById(ctx context.Context, id uuid.UUID) (*[]models.Booking, error){
	var result []models.Booking
	err := m.Where(&models.Booking{Customer_id: id}).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}