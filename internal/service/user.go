package service

import (
	"mio-init/internal/repository"
)

type UserService interface {
}

type userSvc struct {
	userRepo repository.UserRepository
}

func NewUserService() UserService {
	return &userSvc{}
}
