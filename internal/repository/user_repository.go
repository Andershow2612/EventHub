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


func (r *UserRepository) List() ([]entity.UserTeste, error){
	var user []entity.UserTeste

	err := r.DB.Find(&user)
	if err != nil{
		fmt.Println("Users not found")
	}

	return user, nil
}