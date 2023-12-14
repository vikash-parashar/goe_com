package models

import "time"

type User struct {
	ID          uint
	Username    string
	Mobile      string
	Gender      string
	DateOfBirth time.Time
	Email       string
	Password    string
	Cart        Cart
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
