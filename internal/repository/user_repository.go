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

	err := r.DB.Find(&user)
	if err != nil{
		fmt.Println("Users not found")
	}

	return user, nil
}

func (r *UserRepository) ListById(id int) (*entity.User, error){
	var User entity.User

	if err := r.DB.Where("id = ?", id).First(&User).Error; err != nil{
		return nil, err
	}
	return &User, nil
}

func (r *UserRepository) Create(user *entity.User) (*entity.User, error){
	
	if err := r.DB.Create(user).Error; err != nil{
		return nil, err
	}

	return user, nil
}