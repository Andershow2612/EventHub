package repository

import (
	"errors"

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

func (r *EventRepository) Create(event *entity.Event) (*entity.Event, error){
	
	result := r.DB.Create(&event)
	if result.RowsAffected == 0{
		return nil, result.Error
	}

	if result.Error != nil{
		return nil, result.Error
	}

	if err := r.DB.
        Preload("Organizer").
        Preload("Mode").
        Preload("Address").
        Preload("Category").
        First(event, event.ID).Error; err != nil {
        return nil, err
    }

	return event, nil
}

func (r *EventRepository) Delete(eventID, userID int) (string, error) {
    var event entity.Event

    result := r.DB.Debug().First(&event, eventID)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return "", gorm.ErrRecordNotFound
        }
        return "", result.Error
    }

    if event.OrganizerID != userID {
        return "", errors.New("forbidden")
    }

    result = r.DB.Delete(&event)
    if result.Error != nil {
        return "", result.Error
    }

    if result.RowsAffected == 0 {
        return "", errors.New("event not deleted")
    }

    return "event deleted", nil
}
