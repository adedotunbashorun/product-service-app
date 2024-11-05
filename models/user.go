package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-"`       // Store hashed password
	RoleID   uint   `json:"role_id"` // Foreign key to Role
	Role     Role   `json:"role" gorm:"foreignKey:RoleID"`
}
