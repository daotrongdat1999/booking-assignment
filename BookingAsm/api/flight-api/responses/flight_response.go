package responses

import "time"

type FlightResponse struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	From          string    `json:"from"`
	To            string    `json:"to"`
	Date          time.Time `json:"date"`
	Status        string    `json:"status"`
	AvailableSlot int32     `json:"available_slot"`
}
