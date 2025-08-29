package repository

import (
	"eventHub.com/internal/entity"
	"gorm.io/gorm"
)

type EventRepository struct {
	DB *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository{
	return &EventRepository{DB: db}
}

func (r *EventRepository) List() ([]entity.Event, error){
	var event []entity.Event

	result := r.DB.Debug().Preload("Organizer").Preload("Mode").Preload("Address").Preload("Category").Find(&event)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}

	if result.Error != nil{
		return nil, result.Error
	}

	return event, nil
}


func (r *EventRepository) ListById(id int) (*entity.Event, error){
	var event *entity.Event

	result := r.DB.Debug().Preload("Organizer").Preload("Mode").Preload("Address").Preload("Category").Where("id = ?", id).First(&event)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}

	if result.Error != nil{
		return nil, result.Error
	}
	
	return event, nil
}