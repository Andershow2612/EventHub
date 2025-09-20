package repository

import (
	"eventHub.com/internal/entity"
	"gorm.io/gorm"
)

type KindRepository struct {
	DB *gorm.DB
}


func NewKindRepository(db *gorm.DB) *KindRepository{
	return &KindRepository{
		DB: db,
	}
}

func (r *KindRepository) FindById(id uint) (*entity.Kind, error){
	var kind entity.Kind

	 if err := r.DB.Where("id = ?", id).First(&kind).Error; err != nil{
		return nil, err
	 }

	return &kind, nil
}