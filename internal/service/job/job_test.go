package job

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mohammaderm/todoList/internal/dto"
	"github.com/mohammaderm/todoList/internal/models"
	"github.com/mohammaderm/todoList/internal/repository/job/mocks"
	"github.com/mohammaderm/todoList/log"
)

func TestCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	job := models.CreateJob{
		Name:        "homeWork",
		Description: "write homeworks",
		AccountId:   1,
		Status:      false,
	}
	mockRepo := mocks.NewMockJobRepository(mockCtrl)
	mockCache := mocks.NewMockJobCacheInterface(mockCtrl)
	mockRepo.EXPECT().Create(gomock.Any(), job).Return(nil).Times(1)
	mockCache.EXPECT().DeleteAll(gomock.Any()).Return(nil).Times(1)

	log, _ := log.New(&log.Logconfig{})
	jobSerice := NewService(log, mockRepo, mockCache)
	ctx := context.Background()
	err := jobSerice.Create(ctx, dto.CreateJobReq{
		Job: job,
	})
	if err != nil {
		t.Fail()
	}

}
