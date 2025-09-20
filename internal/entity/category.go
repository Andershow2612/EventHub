package entity

import "time"

type Category struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Created_at  time.Time
}

func (Category) TableName() string{
	return "category"
}