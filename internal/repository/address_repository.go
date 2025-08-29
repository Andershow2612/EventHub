package repository

import (
	"eventHub.com/internal/entity"
	"gorm.io/gorm"
)

type AddressRepository struct {
	DB *gorm.DB
}

func NewAddressRepository(db *gorm.DB) *AddressRepository{
	return &AddressRepository{DB: db}
}

func (r *AddressRepository) List() ([]entity.Address, error){
	var address []entity.Address

	err := r.DB.Find(&address).Error
	if err != nil {
		return nil, err
	}

	return address, nil
}

func (r *AddressRepository) ListById(id int) (*entity.Address, error){
	var address *entity.Address

	err := r.DB.Debug().First(&address).Where("id = ?", id).Error
	
	if err != nil {
		return nil, err
	}

	return address, nil
}

func (r *AddressRepository) Create(address *entity.Address) (*entity.Address, error){
	err := r.DB.Create(&address).Error
	if err != nil {
		return nil, err
	}

	return address, nil
}

func (r *AddressRepository) Delete(id int) (string, error){
	result := r.DB.Debug().Where("id = ?", id).Delete(&entity.Address{})
	if result.Error != nil{
		return "The data was not deleted", result.Error
	}

	if result.RowsAffected == 0 {
		return "No record found to delete", nil
	}

	return "Deleted successfully", nil
}