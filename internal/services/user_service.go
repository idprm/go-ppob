package services

import "github.com/idprm/go-ppob/internal/domain/repository"

type UserService struct {
	userRepo repository.IUserRepository
}

func NewUserService(
	userRepo repository.IUserRepository,
) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

type IUserService interface {
}
