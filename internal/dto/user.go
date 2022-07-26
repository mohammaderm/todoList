package dto

import "github.com/mohammaderm/todoList/internal/models"

type (
	CreateUserReq struct {
		Username string `json:"username" db:"username"`
		Email    string `json:"email" db:"email"`
		Password string `json:"password" db:"password"`
	}
	GetByEmailReq struct {
		Email string `json:"email" db:"email"`
	}
	GetByEmailRes struct {
		User *models.User `json:"user" db:"user"`
	}
	GetByUsernameReq struct {
		Username string `json:"username" db:"username"`
	}
	GetByUsernameRes struct {
		User *models.User `json:"user" db:"user"`
	}
	UserLogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)
