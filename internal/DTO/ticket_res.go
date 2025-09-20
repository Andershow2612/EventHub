package dto

type TicketResponse struct {
	ID      uint
	Price   float64
	Name    string
	EventID uint
	Event   EventSummary
	KindID  uint
	Kind    KindReponse
}