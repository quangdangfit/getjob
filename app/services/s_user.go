package services

import (
	"context"

	"github.com/quangdangfit/getjob/app/interfaces"
	"github.com/quangdangfit/getjob/app/models"
	"github.com/quangdangfit/getjob/app/schema"
	"github.com/quangdangfit/getjob/pkg/utils"
)

// UserService handle user logic
type UserService struct {
	repo interfaces.IUserRepository
}

// NewUserService return new IUserService interface
func NewUserService(repo interfaces.IUserRepository) interfaces.IUserService {
	return &UserService{
		repo: repo,
	}
}

// Register handle logic register new user
func (s *UserService) Login(ctx context.Context, params *schema.UserLoginParams) (*models.User, error) {
	user, err := s.repo.GetByEmail(ctx, params.Email)
	if err != nil {
		return nil, err
	}

	ok, err := user.CheckPassword(params.Password)
	if !ok {
		return nil, err
	}

	return user, nil
}

// Register handle logic register new user
func (s *UserService) Register(ctx context.Context, params *schema.UserRegisterParams) (*models.User, error) {
	var user *models.User
	err := utils.Copy(&user, &params)
	if err != nil {
		return nil, err
	}

	err = s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
