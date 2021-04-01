package services

import "github.com/quangdangfit/getjob/app/interfaces"

type UserService struct {
	repo interfaces.IUserRepository
}

func NewUserService(repo interfaces.IUserRepository) interfaces.IUserService {
	return &UserService{
		repo: repo,
	}
}
