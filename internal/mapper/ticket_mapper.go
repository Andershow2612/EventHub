package mapper

import (
	dto "eventHub.com/internal/DTO"
	"eventHub.com/internal/entity"
)

func ToTicketResponse(ticket entity.Ticket) dto.TicketResponse{
	return dto.TicketResponse{
		ID: ticket.ID,
		Name: ticket.Name,
		Price: ticket.Price,
		EventID: ticket.Event.ID,
		Event: EventSummary(&ticket.Event),
	}
}

func ToTicketResponseList(tickets []entity.Ticket) []dto.TicketResponse{

	var responses []dto.TicketResponse

	for _, ticket := range tickets{
		responses = append(responses, ToTicketResponse(ticket))
	}
	return responses
}

func ToTicketEntity(req dto.TicketRequest) entity.Ticket{
	return entity.Ticket{
		Name: req.Name,
		Price: req.Price,
		Sales_start: req.Sales_start,
		EventID: req.EventID,
	}
}