package dto

import "time"

type EventRequest struct {
	Title       string
	Description string
	Start_at    time.Time
	End_at time.Time
	Published_at time.Time
	Created_at time.Time
	Updated_at time.Time
	OrganizerID uint
	ModeID uint
	AddressID uint
	CategoryID uint
}

type EventRequestUpdate struct{
	Title       string
	Description string
	OrganizerID uint
	ModeID 		uint
	AddressID 	uint
	CategoryID 	uint
}