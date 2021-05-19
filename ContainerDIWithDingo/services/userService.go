package services

import "Go/ContainerDIWithDingo/models"

type IUserService interface {
	GetById() (*models.User, error)
}

type userService struct {
	//userRepository repositories.IUserRepository
}

func (s *userService) GetById() (*models.User, error) {
	return &models.User{"fdkasfjdaslfk", "Fulano da Silva"}, nil
}

func NewUserService() IUserService {
	return &userService{}
}
