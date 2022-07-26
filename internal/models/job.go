package models

import "time"

type (
	Job struct {
		Id          uint      `json:"id" db:"id"`
		Name        string    `json:"name" db:"name"`
		Description string    `json:"description" db:"description"`
		AccountId   uint      `json:"accountid" db:"accountid"`
		Status      bool      `json:"status" db:"status"`
		CreatedAt   time.Time `json:"created_at" db:"created_at"`
	}

	CreateJob struct {
		Name        string `json:"name" db:"name"`
		Description string `json:"description" db:"description"`
		AccountId   uint   `json:"-" db:"accountid"`
		Status      bool   `json:"-" db:"status"`
	}
)
