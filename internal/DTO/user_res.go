package dto

import "time"

type UserResponse struct {
	ID         int
	Name       string
	Age        int
	Email      string
	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
	Active int
}