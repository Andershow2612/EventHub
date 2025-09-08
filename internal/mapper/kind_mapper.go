package mapper

import (
	dto "eventHub.com/internal/DTO"
	"eventHub.com/internal/entity"
)

func ToKindResponse(kind entity.Kind) dto.KindReponse{
	return dto.KindReponse{
		ID: kind.ID,
		Name: kind.Name,
	}
}