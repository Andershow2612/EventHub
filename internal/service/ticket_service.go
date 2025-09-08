package service

import (
	"errors"

	"eventHub.com/internal/entity"
	"eventHub.com/internal/repository"
)

type TicketService struct {
	Repo *repository.TicketRepository
	eventRepo  *repository.EventRepository
	kindRepo *repository.KindRepository
}

func NewTicketService(repo repository.TicketRepository, eventRepo repository.EventRepository, kindRepo repository.KindRepository) *TicketService{
	return &TicketService{
		Repo: &repo,
		eventRepo: &eventRepo,
		kindRepo: &kindRepo,
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

	ticketKind, err := s.kindRepo.FindById(ticket.KindID)
	if err != nil{
		return err
	}

	if ticket.KindID == 1 && ticket.Price == 0 {
		return errors.New("ticket must have a value")
	}

	if ticket.KindID == 2 && ticket.Price > 0 {
		return errors.New("tiket must not have a value")
	}

	ticket.EventID = eventID
	ticket.KindID = ticketKind.ID
	return s.Repo.Create(&ticket)
}

func (s *TicketService) GetTicketsByEvent(eventID int) ([]entity.Ticket, error) {
	return s.Repo.FindByEvent(eventID)
}