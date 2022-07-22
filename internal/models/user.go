package models

import "time"

type (
	User struct {
		Id        uint      `json:"id" db:"id"`
		Username  string    `json:"username" db:"username"`
		Email     string    `json:"email" db:"email"`
		Password  string    `json:"password" db:"password"`
		CreatedAt time.Time `json:"created_at" db:"created_at"`
	}
)
