package entity

type Kind struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

func (Kind) TableName() string {
	return "kind"
}