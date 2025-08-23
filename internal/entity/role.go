package entity

import "time"

type Role struct {
	ID         int `gorm:"column:id;primaryKey"`
	Name       string
	Created_at time.Time
}

func (Role) TableName() string{
	return "role"
}