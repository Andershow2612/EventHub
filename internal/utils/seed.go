package utils

import (
	"eventHub.com/internal/entity"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	roles := []entity.Role{
		{ID: 1, Name: "admin"},
		{ID: 2, Name: "organizer"},
		{ID: 3, Name: "user"},
	}

	for _, r := range roles {
		var count int64
		db.Model(&entity.Role{}).Where("id = ?", r.ID).Count(&count)
		if count == 0 {
			db.Create(&r)
		}
	}
}
