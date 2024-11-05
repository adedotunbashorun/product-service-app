package models

type Role struct {
	BaseModel
	Name        string `json:"name" gorm:"unique;not null"`
	Description string `json:"description"`
}

type CreateRoleInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// UpdateRoleInput defines the input structure for updating a role.
type UpdateRoleInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}
