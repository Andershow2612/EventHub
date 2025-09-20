package entity

import "time"

type Event struct {
	ID           uint `gorm:"primaryKey"`
	Title        string
	Description  string
	Banner       []byte
	Start_at     time.Time
	End_at time.Time
	Published_at time.Time
	Created_at time.Time
	Updated_at time.Time


	OrganizerID uint `gorm:"column:organizer_id"`
	Organizer User `gorm:"foreignKey:OrganizerID;references:ID"`

	ModeID uint `gorm:"column:mode_id"`
	Mode Mode `gorm:"foreignKey:ModeID;references:ID"`

	AddressID uint `gorm:"column:address_id"`
	Address Address `gorm:"foreignKey:AddressID;references:ID"`
	
	CategoryID uint `gorm:"column:category_id"`
	Category Category `gorm:"foreignKey:CategoryID;references:ID"`
}

func (Event) TableName() string{
	return "event"
}