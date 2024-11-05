package seeder

import (
	"user-management-mysql/models"

	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) error {
	roles := []models.Role{
		{Name: "Admin"},
		{Name: "User"},
	}

	for _, role := range roles {
		var existingRole models.Role
		db.Where("name = ?", role.Name).First(&existingRole)
		if existingRole.ID == 0 {
			if err := db.Create(&role).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
