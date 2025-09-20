package entity

import "time"

type Ticket struct {
	ID          uint `gorm:"primaryKey"`
	Price       float64
	Name        string
	Sales_start time.Time
	Sales_end time.Time
	Quantity_max int
	Quantity_sold int

	EventID     uint 
	Event       Event `gorm:"foreignKey:EventID;references:ID"`

	KindID uint
	Kind Kind `gorm:"foreignKey:KindID;references:ID"`
}

func (Ticket) TableName() string{
	return "event_ticket"
}