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

func (s *AddressService) ListAddrresID(id int) (*entity.Address, error){
	return s.Repo.ListById(id)
}

func (s *AddressService) CreateAddress(address *entity.Address) (*entity.Address, error){
	return s.Repo.Create(address)
}

func (s *AddressService) DeleteAddress(id int) (string, error){
	return s.Repo.Delete(id)
}