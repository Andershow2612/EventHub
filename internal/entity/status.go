package entity

type Status struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func (Status) TableName() string {
	return "status"
}