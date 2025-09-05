package repository

import (
	"eventHub.com/internal/entity"
	"gorm.io/gorm"
)

type TicketRepository struct {
	DB *gorm.DB
}

func NewTicketRepository(db *gorm.DB) *TicketRepository{
	return &TicketRepository{DB: db}
}

func (r *TicketRepository) List() ([]entity.Ticket, error){
	var ticket []entity.Ticket

	result := r.DB.Debug().Preload("Event.Address").Find(&ticket)
	
	if result.Error != nil{
		return nil, result.Error
	}

	if result.RowsAffected == 0{
		return nil, result.Error
	}

	return ticket, nil
}

func (r *TicketRepository) Create(ticket *entity.Ticket) error {
	
	if err := r.DB.Create(ticket).Error; err != nil{
		return err
	}
	return nil
}

func (r *TicketRepository) FindByEvent(eventID int) ([]entity.Ticket, error) {
	var tickets []entity.Ticket
	
	err := r.DB.Preload("Event").Preload("Event.Address").Where("event_id = ?", eventID).Find(&tickets).Error
	return tickets, err
}