package repositories

import (
	"product-service-app/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) FindByEmail(email string) (models.User, error) {
	var user models.User
	err := repo.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

// GetUserWithRole finds a user by ID and retrieves their associated role
func (repo *UserRepository) GetUserWithRole(userID uint) (*models.User, error) {
	var user models.User
	err := repo.DB.Preload("Role").First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) FindById(id string) (models.User, error) {
	var user models.User
	err := repo.DB.Where("id = ?", id).First(&user).Error
	return user, err
}
