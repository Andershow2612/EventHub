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

func (s *UserService) ListUser() ([]entity.UserTeste, error){
	return s.Repo.List()
}
