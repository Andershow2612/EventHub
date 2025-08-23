package mapper

import (
	dto "eventHub.com/internal/DTO"
	"eventHub.com/internal/entity"
)

func ToRoleResponse(role entity.Role) dto.RoleResponse{
	return dto.RoleResponse{
		ID: role.ID,
		Name: role.Name,
	}
}
