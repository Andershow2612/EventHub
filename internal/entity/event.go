package entity

import "time"

type Event struct {
	ID           int
	Title        string
	Description  string
	Banner       []byte
	Start_at     time.Time
	End_at time.Time
	Published_at time.Time
	Created_at time.Time
	Updated_at time.Time



	Organizer_id int
	Mode_id int
	Address_id int 
	Category_id int
	
}