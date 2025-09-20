package entity

type Mode struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func (Mode) TableName() string {
	return "mode"
}