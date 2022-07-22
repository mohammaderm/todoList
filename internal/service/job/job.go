package job

import (
	"context"

	"github.com/mohammaderm/todoList/internal/repository/job"
	"honnef.co/go/tools/lintcmd/cache"
	"github.com/mohammaderm/todoList/log"
)



type (
	service struct{
		logger log.Logger
		jobrepository job.JobRepository
		cache JobCache
	}

	JobServices interface {
		Create (ctx context.Context, req )
		GetAll (ctx context.Context, req )
		Delete (ctx context.Context, req )
		Update (ctx context.Context, req )
	}
)


func NewService(logger log.Logger, jobrepository job.JobRepository, cache JobCache)  JobService{
	return &service{
		logger:        logger,
		jobrepository: jobrepository,
		cache:         cache,
	}
}

