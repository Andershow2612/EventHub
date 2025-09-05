package service

import (
	"eventHub.com/internal/entity"
	"eventHub.com/internal/repository"
)

type TicketService struct {
	Repo *repository.TicketRepository
	eventRepo  *repository.EventRepository
}

func NewTicketService(repo repository.TicketRepository, eventRepo repository.EventRepository) *TicketService{
	return &TicketService{
		Repo: &repo,
		eventRepo: &eventRepo,
	}
}

func (s *TicketService) List() ([]entity.Ticket, error) {
	return s.Repo.List()
}

func (s *TicketService) CreateTicket(eventID int, ticket entity.Ticket) error{
	_, err := s.eventRepo.ListById(eventID)
	if err != nil{
		return err
	}

	ticket.EventID = eventID
	return s.Repo.Create(&ticket)
}

func (s *TicketService) GetTicketsByEvent(eventID int) ([]entity.Ticket, error) {
	return s.Repo.FindByEvent(eventID)
}