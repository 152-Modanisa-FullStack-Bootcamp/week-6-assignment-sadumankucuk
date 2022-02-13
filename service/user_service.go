package service

import (
	"wallet/model"
	"wallet/repository"
)

type IUserService interface {
	CreateUser(username string) string
	FindAllUser() model.UserResponse
	FindUser(username string) (*model.User, error)
	UpdateBalance(username string, b model.UserRequest) (*model.UserRequest, error)
}

type UserService struct {
	Repository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) IUserService {
	return &UserService{Repository: repository}
}

func (s *UserService) CreateUser(username string) string {
	message := s.Repository.CreateUser(username)
	return message
}

func (s *UserService) FindAllUser() model.UserResponse {
	return s.Repository.FindAllUser()
}

func (s *UserService) FindUser(username string) (*model.User, error) {
	return s.Repository.FindUser(username)
}

func (s *UserService) UpdateBalance(username string, b model.UserRequest) (*model.UserRequest, error) {
	return s.Repository.UpdateBalance(username, b)
}
