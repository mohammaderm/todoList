package dto

import "github.com/mohammaderm/todoList/internal/models"

type (
	CreateJobReq struct {
		Job models.CreateJob `json:"job"`
	}
	GetAllJobRes struct {
		Jobs *[]models.Job `json:"jobs"`
	}
	GetAllJobReq struct {
		AccountId uint `json:"accountid"`
		Offset    int  `json:"offset"`
	}
	DeleteJobReq struct {
		Id        uint `json:"id"`
		AccountId uint `json:"accountid"`
	}
	UpdateJob struct {
		Id        uint `json:"id"`
		AccountId uint `json:"accountid"`
	}
)
