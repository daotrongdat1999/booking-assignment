package responses

import (
	cr "BookingAsm/api/customer-api/responses"
	fr "BookingAsm/api/flight-api/responses"
	"time"
)

type BookingResponse struct {
	ID          string    `json:"id"`
	Customer_id string    `json:"customer_id"`
	Flight_id   string    `json:"flight_id"`
	Code        string    `json:"code"`
	Status      string    `json:"status"`
	Booked_date time.Time `json:"booked_date"`
}

type SearchBookingResponse struct {
	BookingResponse
	CustomerInfor cr.CustomerResponse
	FlightInfor   fr.FlightResponse
}
