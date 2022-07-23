package job

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mohammaderm/todoList/internal/models"
	"github.com/mohammaderm/todoList/log"
)

type (
	jobCache struct {
		logger log.Logger
		redis  *redis.Client
	}
	JobCacheInterface interface {
		SetAll(ctx context.Context, key string, value *[]models.Job, ttl time.Duration) error
		GetAll(ctx context.Context, key string) (*[]models.Job, error)
		DeleteAll(ctx context.Context) error
	}
)

func NewJObCache(logger log.Logger, redis *redis.Client) JobCacheInterface {
	return &jobCache{
		logger: logger,
		redis:  redis,
	}
}

func (c jobCache) SetAll(ctx context.Context, key string, value *[]models.Job, ttl time.Duration) error {
	marshalValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = c.redis.Set(ctx, key, marshalValue, ttl).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c jobCache) GetAll(ctx context.Context, key string) (*[]models.Job, error) {
	value, err := c.redis.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var job *[]models.Job
	err = json.Unmarshal([]byte(value), &job)
	if err != nil {
		return nil, err
	}
	return job, nil
}

func (c jobCache) DeleteAll(ctx context.Context) error {
	err := c.redis.FlushDB(ctx).Err()
	if err != nil {
		return err
	}
	return nil
}
