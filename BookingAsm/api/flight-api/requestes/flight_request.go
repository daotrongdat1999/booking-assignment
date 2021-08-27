package requestes

import "time"

type CreatFlightRequest struct {
	Name          string    `json:"name"`
	From          string    `json:"from"`
	To            string    `json:"to"`
	Date          time.Time `json:"date"`
	Status        string    `json:"status"`
	AvailableSlot int32     `json:"available_slot"`
}

//search flight by id
type SearchFlightRequest struct {
	ID string `json:"id" binding:"required"`
}

type UpdateFlightRequest struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	From          string    `json:"from"`
	To            string    `json:"to"`
	Date          time.Time `json:"date"`
	Status        string    `json:"status"`
	AvailableSlot int32     `json:"available_slot"`
}
