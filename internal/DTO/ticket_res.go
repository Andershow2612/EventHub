package dto

type TicketResponse struct {
	ID      int
	Price   float64
	Name    string
	EventID int
	Event   EventSummary
}