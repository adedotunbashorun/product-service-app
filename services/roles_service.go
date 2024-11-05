package services

import (
	"product-service-app/models"
	"product-service-app/repositories"
)

type RoleService struct {
	RoleRepo *repositories.RoleRepository
}

func NewRoleService(roleRepo *repositories.RoleRepository) *RoleService {
	return &RoleService{RoleRepo: roleRepo}
}

func (s *RoleService) CreateRole(input models.CreateRoleInput) (models.Role, error) {
	role := models.Role{
		Name:        input.Name,
		Description: input.Description,
	}
	err := s.RoleRepo.Create(&role)
	return role, err
}

func (s *RoleService) GetRoles() ([]models.Role, error) {
	return s.RoleRepo.GetAll()
}

func (s *RoleService) UpdateRole(id string, input models.UpdateRoleInput) (models.Role, error) {
	role, err := s.RoleRepo.GetByID(id)
	if err != nil {
		return role, err
	}

	if input.Name != nil {
		role.Name = *input.Name
	}
	if input.Description != nil {
		role.Description = *input.Description
	}

	err = s.RoleRepo.Update(&role)
	return role, err
}

func (s *RoleService) DeleteRole(id string) error {
	return s.RoleRepo.Delete(id)
}
