package mapper

import (
	dto "eventHub.com/internal/DTO"
	"eventHub.com/internal/entity"
)

func ToAddressResponse(address *entity.Address) dto.AddressResponse{
	return dto.AddressResponse{
		Street: address.Street,
		HouseNumber: address.HouseNumber,
		City: address.City,
		State: address.State,
		Country: address.Country,
		Complement: address.Complement,
	}
}

func ToAddressResponseList(addresses []entity.Address) []dto.AddressResponse{
	var responses []dto.AddressResponse

	for _, address := range addresses{
		responses = append(responses, ToAddressResponse(&address))
	}
	return responses
}

func ToAddressEntity(req dto.AddressRequest) *entity.Address{
	return &entity.Address{
		City: req.City,
		Street: req.Street,
		HouseNumber: req.HouseNumber,
		State: req.State,
		Country: req.Country,
		Complement: req.Complement,
	}
}