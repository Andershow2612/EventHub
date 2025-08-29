package dto

type AddressRequest struct {
	City        string
	Street      string
	HouseNumber int
	State       string
	Country     string
	Complement  string
}