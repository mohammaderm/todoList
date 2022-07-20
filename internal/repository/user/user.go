package user

import (
	"context"

	"github.com/mohammaderm/todoList/internal/repository/models"
	"github.com/mohammaderm/todoList/log"

	"github.com/jmoiron/sqlx"
)

type (
	repository struct {
		logger log.Logger
		db     *sqlx.DB
	}

	UserRepository interface {
		// user interfaces
		Create(ctx context.Context, email, username, password string) error
		GetbyEmail(ctx context.Context, email string) (*models.User, error)
		GetByUserName(ctx context.Context, username string) (*models.User, error)
	}
)

func NewRepository(con *sqlx.DB, logger log.Logger) UserRepository {
	return &repository{
		logger: logger,
		db:     con,
	}
}

func (r *repository) Create(ctx context.Context, email, username, password string) error {
	_, err := r.db.ExecContext(ctx, CreateUser, username, email, password)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetbyEmail(ctx context.Context, email string) (*models.User, error) {
	var result models.User
	err := r.db.GetContext(ctx, &result, GetUserbyEmail, email)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repository) GetByUserName(ctx context.Context, username string) (*models.User, error) {
	var result models.User
	err := r.db.GetContext(ctx, &result, GetUserbyusername, username)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
