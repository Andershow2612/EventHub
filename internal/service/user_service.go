package service

import (
	"eventHub.com/internal/entity"
	"eventHub.com/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService{
	return &UserService{Repo: &repo}
}

func (s *UserService) ListUser() ([]entity.User, error){
	return s.Repo.List()
}

func (s *UserService) ListAUser(id int) (*entity.User, error){
	return s.Repo.ListById(id)
}

func (s *UserService) CreateUser(user *entity.User) (*entity.User, error){
	return s.Repo.Create(user)
}
