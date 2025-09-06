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
	OrganizerID int
	ModeID int
	AddressID int
	CategoryID int
}

type EventRequestUpdate struct{
	Title       string
	Description string
	OrganizerID int
	ModeID 		int
	AddressID 	int
	CategoryID 	int
}