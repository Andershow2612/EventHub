package service

import (
	"eventHub.com/internal/entity"
	"eventHub.com/internal/repository"
)

type AddressService struct {
	Repo *repository.AddressRepository
}

func NewAddressService(repo repository.AddressRepository) *AddressService{
	return &AddressService{Repo: &repo}
}

func (s *AddressService) ListAddrres() ([]entity.Address, error){
	return s.Repo.List()
}