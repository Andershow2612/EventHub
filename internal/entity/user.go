package entity

import "time"

type User struct {
	ID          int
	Name        string
	Age         int
	Email       string
	Password    string
	Role_id     int
	Profile_img []byte
	Created_at  time.Time
	Updated_at time.Time
	Deleted_at time.Time
	Active int
	Verification_code int
}

func (User) TableName() string {
	return "users"
}