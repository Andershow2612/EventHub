package entity

import "time"

type User struct {
	ID                int       
	Name              string    
	Age               int       
	Email             string    
	Password          string    
	ProfileImg        []byte    
	CreatedAt         time.Time 
	UpdatedAt         time.Time 
	DeletedAt         time.Time 
	Active            int       
	VerificationCode  int       

	RoleID     int `gorm:"column:role_id"`
	Role Role `gorm:"foreignKey:RoleID;references:ID"`
}

func (User) TableName() string {
	return "users"
}