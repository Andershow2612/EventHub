package entity

type Mode struct {
	ID   int `gorm:"primaryKey"`
	Name string
}