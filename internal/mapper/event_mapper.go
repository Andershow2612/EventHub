package mapper

import (
	dto "eventHub.com/internal/DTO"
	"eventHub.com/internal/entity"
)

func ToEventResponse(event *entity.Event) dto.EventResponse{
	return dto.EventResponse{
		ID: event.ID,
		Title: event.Title,
		Description: event.Description,
		Start_at: event.Start_at,
		End_at: event.End_at,
		Published_at: event.Published_at,
		Created_at: event.Created_at,
		Updated_at: event.Updated_at,
		Organizer: ToUserResponse(&event.Organizer),
		Mode: ToModeResponse(&event.Mode),
		Address: ToAddressResponse(&event.Address),
		Category: ToCategoryResponse(event.Category),
	}
}

func ToEventResponseList(events []entity.Event) []dto.EventResponse{
	var responses []dto.EventResponse

	for _, event := range events{
		responses = append(responses, ToEventResponse(&event))
	}
	return responses
}

func ToEventEntity(event dto.EventRequest) *entity.Event{
	return &entity.Event{
		Title: event.Title,
		Description: event.Description,
		Start_at: event.Start_at,
		End_at: event.End_at,
		Published_at: event.Published_at,
		Created_at: event.Created_at,
		OrganizerID: event.OrganizerID,
		ModeID: event.ModeID,
		AddressID: event.AddressID,
		CategoryID: event.CategoryID,
	}
}

func ToEventUpdate(event dto.EventRequestUpdate) entity.Event {
	return entity.Event{
		Title:       event.Title,
		Description: event.Description,
		OrganizerID: event.OrganizerID,
		ModeID:      event.ModeID,
		AddressID:   event.AddressID,
		CategoryID:  event.CategoryID,
	}
}

func EventSummary(event *entity.Event) dto.EventSummary{
	return dto.EventSummary{
		ID: event.ID,
		Title: event.Title,
		Description: event.Description,
		Start_at: event.Start_at,
		Address: ToAddressResponse(&event.Address),
	}
}