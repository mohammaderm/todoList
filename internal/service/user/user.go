package user

import (
	"context"

	"github.com/mohammaderm/todoList/internal/dto"
	"github.com/mohammaderm/todoList/internal/repository/user"
	"github.com/mohammaderm/todoList/log"
)

type (
	Service struct {
		logger         log.Logger
		userRepository user.UserRepository
	}
	UserServiceContracts interface {
		// user services
		Create(ctx context.Context, req dto.CreateUserReq) error
		GetbyEmail(ctx context.Context, req dto.GetByEmailReq) (dto.GetByEmailRes, error)
		GetByUserName(ctx context.Context, req dto.GetByUsernameReq) (dto.GetByUsernameRes, error)
	}
)

func NewService(logger log.Logger, userrepository user.UserRepository) UserServiceContracts {
	return &Service{
		logger:         logger,
		userRepository: userrepository,
	}
}

func (s *Service) Create(ctx context.Context, req dto.CreateUserReq) error {
	err := s.userRepository.Create(ctx, req.Email, req.Username, req.Password)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetbyEmail(ctx context.Context, req dto.GetByEmailReq) (dto.GetByEmailRes, error) {
	user, err := s.userRepository.GetbyEmail(ctx, req.Email)
	if err != nil {
		return dto.GetByEmailRes{}, err
	}
	return dto.GetByEmailRes{
		User: user,
	}, nil
}

func (s *Service) GetByUserName(ctx context.Context, req dto.GetByUsernameReq) (dto.GetByUsernameRes, error) {
	user, err := s.userRepository.GetByUserName(ctx, req.Username)
	if err != nil {
		return dto.GetByUsernameRes{}, err
	}
	return dto.GetByUsernameRes{
		User: user,
	}, nil
}
