package job

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/mohammaderm/todoList/internal/dto"
	"github.com/mohammaderm/todoList/internal/repository/job"
	"github.com/mohammaderm/todoList/log"
)

type (
	Service struct {
		logger        log.Logger
		jobrepository job.JobRepository
		cache         JobCacheInterface
	}

	JobServices interface {
		Create(ctx context.Context, req dto.CreateJobReq) error
		GetAll(ctx context.Context, req dto.GetAllJobReq) (dto.GetAllJobRes, error)
		Delete(ctx context.Context, req dto.DeleteJobReq) error
		Update(ctx context.Context, req dto.UpdateJob) error
	}
)

func NewService(logger log.Logger, jobrepository job.JobRepository, cache JobCacheInterface) JobServices {
	return &Service{
		logger:        logger,
		jobrepository: jobrepository,
		cache:         cache,
	}
}

func (s *Service) Create(ctx context.Context, req dto.CreateJobReq) error {
	err := s.jobrepository.Create(ctx, req.Job)
	if err != nil {
		return err
	}
	s.cache.DeleteAll(ctx)
	return nil
}

func (s *Service) GetAll(ctx context.Context, req dto.GetAllJobReq) (dto.GetAllJobRes, error) {
	result, err := s.cache.GetAll(ctx, fmt.Sprint(req.Offset))
	if err == redis.Nil {
		s.logger.Warning("can not found value in cache", map[string]interface{}{
			"err": err.Error(),
		})
		result, err = s.jobrepository.GetAll(ctx, req.AccountId, req.Offset)
		if err != nil {
			return dto.GetAllJobRes{}, err
		}
		return dto.GetAllJobRes{
			Jobs: result,
		}, nil
	}
	println("cache hit")
	return dto.GetAllJobRes{
		Jobs: result,
	}, nil
}

func (s *Service) Delete(ctx context.Context, req dto.DeleteJobReq) error {
	err := s.jobrepository.Delete(ctx, req.Id, req.AccountId)
	if err != nil {
		return err
	}
	err = s.cache.DeleteAll(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(ctx context.Context, req dto.UpdateJob) error {
	err := s.jobrepository.Update(ctx, req.Id, req.AccountId)
	if err != nil {
		return err
	}
	err = s.cache.DeleteAll(ctx)
	if err != nil {
		return err
	}
	return nil
}
