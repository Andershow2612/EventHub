package mapper

import (
	dto "eventHub.com/internal/DTO"
	"eventHub.com/internal/entity"
)

func ToCategoryResponse(category entity.Category) dto.CategoryResponse{
	return dto.CategoryResponse{
		ID: category.ID,
		Name: category.Name,
		Description: category.Description,
	}
}