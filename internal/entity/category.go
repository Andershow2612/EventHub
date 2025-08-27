package entity

import "time"

type Category struct {
	ID          int
	Name        string
	Description string
	Created_at  time.Time
}