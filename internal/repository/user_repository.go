package repository

import (
	"fmt"

	"eventHub.com/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db * gorm.DB) *UserRepository{
	return &UserRepository{DB: db}
}


func (r *UserRepository) List() ([]entity.User, error){
	var user []entity.User

	err := r.DB.Debug().Preload("Role").Find(&user).Error
	if err == nil{
		fmt.Println("Users not found")
	}

	if err != nil{
		return  nil, err
	}

	return user, nil
}

func (r *UserRepository) ListById(id int) (*entity.User, error) {
	var user entity.User

	if err := r.DB.Debug().Preload("Role").Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) ListByEmail(email string) (*entity.User, error){
	var user entity.User

	if err := r.DB.Debug().Preload("Role").Where("email = ?", email).Find(&user).Error; err != nil{
		return nil, err
	}
	
	return &user, nil
}

func (r *UserRepository) Create(user *entity.User) (*entity.User, error){
	
	if err := r.DB.Create(user).Error; err != nil{
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) UpdateSelectedFields(user *entity.User, fields []string) (*entity.User, error) {
    if err := r.DB.Model(user).Select(fields).Updates(user).Error; err != nil {
        return nil, err
    }
    return user, nil
}

func (r *UserRepository) Update(user *entity.User) error{
	return r.DB.Save(user).Error
}