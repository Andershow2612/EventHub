package dto

import "time"

type EventResponse struct {
	ID          int
	Title       string
	Description string
	Start_at    time.Time
	End_at time.Time
	Published_at time.Time
	Created_at time.Time
	Updated_at time.Time

	Organizer UserResponse
	Mode ModeResponse
	Address AddressResponse
	Category CategoryResponse
}