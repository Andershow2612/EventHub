package service

import (
	"time"

	"eventHub.com/internal/entity"
	"eventHub.com/internal/repository"
)

type EventService struct {
	Repo repository.EventRepository
}

func NewEventService(repo repository.EventRepository) *EventService{
	return &EventService{
		Repo: repo,
	}
}

func (s *EventService) ListEvents() ([]entity.Event, error){
	return s.Repo.List()
}

func (s *EventService) ListEventID(id int) (*entity.Event, error){
	return s.Repo.ListById(id)
}

func(s *EventService) Update(id int, event *entity.Event) (*entity.Event, error) {
	
	exist, err := s.Repo.ListById(id)
	if err != nil {
		return nil, err
	}

	exist.Title = event.Title
	exist.Description = event.Description
	exist.Updated_at = time.Now()
	exist.OrganizerID = event.OrganizerID
	exist.ModeID = event.ModeID
	exist.AddressID = event.AddressID
	exist.CategoryID = event.CategoryID
	
	return s.Repo.Update(exist)
}

func (s *EventService) CreateEvent(event *entity.Event) (*entity.Event, error){

	createdEvent, err := s.Repo.Create(event)
	if err != nil{
		return nil, err
	}

	return createdEvent, nil
}

func (s *EventService) DeleteEvent(eventID, userID int) (string, error){
	return s.Repo.Delete(eventID, userID)
}