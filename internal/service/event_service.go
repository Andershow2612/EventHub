package service

import (
	"eventHub.com/internal/entity"
	"eventHub.com/internal/repository"
)

type EventService struct {
	Repo *repository.EventRepository
}

func NewEventService(repo repository.EventRepository) *EventService{
	return &EventService{Repo: &repo}
}

func (s *EventService) ListEvents() ([]entity.Event, error){
	return s.Repo.List()
}

func (s *EventService) ListEventID(id int) (*entity.Event, error){
	return s.Repo.ListById(id)
}