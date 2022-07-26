package job

import (
	"context"
	"errors"

	"github.com/mohammaderm/todoList/internal/models"
	"github.com/mohammaderm/todoList/log"

	"github.com/jmoiron/sqlx"
)

var ErrorNotFound = errors.New("can not founf any job with this id")

type (
	repository struct {
		logger log.Logger
		db     *sqlx.DB
	}

	JobRepository interface {
		// job interfaces
		Create(ctx context.Context, job models.CreateJob) error
		GetAll(ctx context.Context, accountid uint, offset int) (*[]models.Job, error)
		Delete(ctx context.Context, jobid, accountid uint) error
		Update(ctx context.Context, jobid, accountid uint) error
	}
)

func NewRepository(con *sqlx.DB, logger log.Logger) JobRepository {
	return &repository{
		logger: logger,
		db:     con,
	}
}

// create new job
func (r *repository) Create(ctx context.Context, job models.CreateJob) error {
	_, err := r.db.ExecContext(ctx, CreateJob, job.Name, job.Description, job.AccountId)
	if err != nil {
		return err
	}
	return nil
}

// get all account job
func (r *repository) GetAll(ctx context.Context, accountid uint, offset int) (*[]models.Job, error) {
	var result []models.Job
	err := r.db.SelectContext(ctx, &result, GetAllJob, accountid, limit, offset)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repository) Delete(ctx context.Context, jobid, accountid uint) error {
	result, err := r.db.ExecContext(ctx, DeleteJob, jobid, accountid)
	if err != nil {
		return err
	}
	rowEfected, _ := result.RowsAffected()
	if rowEfected == 0 {
		return ErrorNotFound
	}
	return nil
}

func (r *repository) Update(ctx context.Context, jobid, accountid uint) error {
	result, err := r.db.ExecContext(ctx, UpdateJob, jobid, accountid)
	if err != nil {
		return err
	}
	rowEfected, _ := result.RowsAffected()
	if rowEfected == 0 {
		return ErrorNotFound
	}
	return nil
}
