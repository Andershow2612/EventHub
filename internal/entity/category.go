package entity

import "time"

type Category struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Description string
	Created_at  time.Time
}