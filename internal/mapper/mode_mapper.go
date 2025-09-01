package mapper

import (
	dto "eventHub.com/internal/DTO"
	"eventHub.com/internal/entity"
)

func ToModeResponse(mode *entity.Mode) dto.ModeResponse{
	return dto.ModeResponse{
		ID: mode.ID,
		Name: mode.Name,
	}
}