package entity

type Status struct {
	ID   int `gorm:"primaryKey"`
	Name string
}