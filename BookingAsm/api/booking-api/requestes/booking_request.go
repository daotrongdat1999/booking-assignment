package requestes

import "time"

type CreatBookingRequest struct {
	Customer_id string    `json:"customer_id"`
	Flight_id   string    `json:"flight_id"`
	Code        string    `json:"code"`
	Status      string    `json:"status"`
	Booked_date time.Time `json:"booked_date"`
}

//search booking by code
type SearchBookingRequest struct {
	Code string `json:"code" binding:"required"`
}

type CancelBookingRequest struct {
	Code   string `json:"code"`
	Status string `json:"status"`
}

// type ViewAllBookingById struct{
// 	Customer_id string `json:"code" binding:"customer_id"`
// }
