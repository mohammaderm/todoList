package job

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/mohammaderm/todoList/internal/dto"
	"github.com/mohammaderm/todoList/internal/models"
	"github.com/mohammaderm/todoList/internal/repository/job/mocks"
	"github.com/mohammaderm/todoList/log"
	"github.com/stretchr/testify/require"
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
	mockRepo.EXPECT().Create(gomock.Any(), job).Return(nil).Times(1)

	mockCache := mocks.NewMockJobCacheInterface(mockCtrl)
	mockCache.EXPECT().DeleteAll(gomock.Any()).Return(nil).Times(1)

	log, _ := log.New(&log.Logconfig{})
	jobService := NewService(log, mockRepo, mockCache)

	ctx := context.Background()
	err := jobService.Create(ctx, dto.CreateJobReq{
		Job: job,
	})

	require.NoError(t, err)
}

func TestGetAll(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	job1 := models.Job{
		Id:          1,
		Name:        "homeWork",
		Description: "write homeworks",
		AccountId:   1,
		Status:      false,
		CreatedAt:   time.Time{},
	}

	job2 := models.Job{
		Id:          2,
		Name:        "homeWork",
		Description: "write homeworks",
		AccountId:   1,
		Status:      false,
		CreatedAt:   time.Time{},
	}

	jobs := []models.Job{job1, job2}

	req := dto.GetAllJobReq{
		AccountId: 1,
		Offset:    1,
	}

	mockRepo := mocks.NewMockJobRepository(mockCtrl)
	mockRepo.EXPECT().GetAll(gomock.Any(), gomock.Eq(req.AccountId), gomock.Eq(req.Offset)).Return(&jobs, nil).Times(1)

	mockCache := mocks.NewMockJobCacheInterface(mockCtrl)
	mockCache.EXPECT().GetAll(gomock.Any(), "1").Return(&jobs, nil).Times(1)

	log, _ := log.New(&log.Logconfig{})
	jobService := NewService(log, mockRepo, mockCache)

	ctx := context.Background()
	result, err := jobService.GetAll(ctx, req)

	require.NoError(t, err)
	require.Equal(t, dto.GetAllJobRes{
		Jobs: &jobs,
	}, result)

}

func TestDelete(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	req := dto.DeleteJobReq{
		Id:        1,
		AccountId: 1,
	}
	mockRepo := mocks.NewMockJobRepository(mockCtrl)
	mockRepo.EXPECT().Delete(gomock.Any(), gomock.Eq(req.Id), gomock.Eq(req.AccountId)).Return(nil).Times(1)

	mockCache := mocks.NewMockJobCacheInterface(mockCtrl)
	mockCache.EXPECT().DeleteAll(gomock.Any()).Return(nil).Times(1)

	log, _ := log.New(&log.Logconfig{})
	jobService := NewService(log, mockRepo, mockCache)

	ctx := context.Background()
	err := jobService.Delete(ctx, req)

	require.NoError(t, err)
}
