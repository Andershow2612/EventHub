package entity

import "time"

type Ticket struct {
	ID          int `gorm:"primaryKey"`
	Price       float64
	Name        string
	Sales_start time.Time
	Sales_end time.Time
	Quantity_max int
	Quantity_sold int

	EventID     int 
	Event       Event `gorm:"foreignKey:EventID;references:ID"`
}

func (Ticket) TableName() string{
	return "event_ticket"
}