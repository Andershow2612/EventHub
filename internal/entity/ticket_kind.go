package entity

type TicketKind struct {
	ID   uint
	Name string
}

func (TicketKind) TableName() string {
	return "kind"
}