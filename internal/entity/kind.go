package entity

type Kind struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func (Kind) TableName() string {
	return "kind"
}