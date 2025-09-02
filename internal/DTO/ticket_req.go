package dto

import "time"

type TicketRequest struct {
	Price       float64
	Name        string
	Sales_start time.Time
	EventID     int 
	Quantity_max int
}