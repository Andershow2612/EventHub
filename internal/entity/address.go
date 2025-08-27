package entity

type Address struct {
	ID          int `gorm:"primaryKey"`
	City        string
	Street      string
	HouseNumber int
	State       string
	Country     string
	Complement  string
}