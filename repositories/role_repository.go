package repositories

import (
	"product-service-app/models"

	"gorm.io/gorm"
)

type RoleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{DB: db}
}

func (r *RoleRepository) Create(role *models.Role) error {
	return r.DB.Create(role).Error
}

func (r *RoleRepository) GetAll() ([]models.Role, error) {
	var roles []models.Role
	err := r.DB.Find(&roles).Error
	return roles, err
}

func (r *RoleRepository) GetByID(id string) (models.Role, error) {
	var role models.Role
	err := r.DB.First(&role, "id = ?", id).Error // Change ID type to string
	return role, err
}

func (r *RoleRepository) Update(role *models.Role) error {
	return r.DB.Save(role).Error
}

func (r *RoleRepository) Delete(id string) error {
	return r.DB.Delete(&models.Role{}, "id = ?", id).Error // Change ID type to string
}

func (r *RoleRepository) FindByName(name string) (*models.Role, error) {
	var role models.Role
	if err := r.DB.Where("name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepository) FindById(id uint) (*models.Role, error) {
	var role models.Role
	if err := r.DB.Where("id = ?", id).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
