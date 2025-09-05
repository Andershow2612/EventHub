package entity

type TicketKind struct {
	ID   string
	Name string
}

func (TicketKind) TableName() string {
	return "kind"
}