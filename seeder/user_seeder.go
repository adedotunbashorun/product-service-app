package seeder

import (
	"fmt"
	"product-service-app/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SeedUsers will seed the User table after roles have been seeded
func SeedUsers(db *gorm.DB) error {
	// Fetch roles to associate with the users
	var adminRole models.Role
	var userRole models.Role
	if err := db.Where("name = ?", "Admin").First(&adminRole).Error; err != nil {
		return fmt.Errorf("could not find Admin role: %w", err)
	}
	if err := db.Where("name = ?", "User").First(&userRole).Error; err != nil {
		return fmt.Errorf("could not find User role: %w", err)
	}

	// Define users to be seeded
	users := []models.User{
		{
			Username: "admin_user",
			Email:    "admin@example.com",
			Password: "adminpassword", // Will hash this password
			RoleID:   adminRole.ID,
		},
		{
			Username: "normal_user",
			Email:    "user@example.com",
			Password: "userpassword", // Will hash this password
			RoleID:   userRole.ID,
		},
	}

	// Iterate through each user and hash the password before saving
	for _, user := range users {
		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("could not hash password for user %s: %w", user.Username, err)
		}
		user.Password = string(hashedPassword)

		// Check if the user already exists
		var existingUser models.User
		db.Where("email = ?", user.Email).First(&existingUser)
		if existingUser.ID == 0 { // If no user exists, create one
			if err := db.Create(&user).Error; err != nil {
				return fmt.Errorf("could not create user %s: %w", user.Username, err)
			}
		}
	}

	return nil
}
