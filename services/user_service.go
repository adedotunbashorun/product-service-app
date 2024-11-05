package services

import (
	"errors"
	"user-management-mysql/models"
	"user-management-mysql/repositories"
	"user-management-mysql/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo *repositories.UserRepository
	RoleRepo *repositories.RoleRepository
}

func NewUserService(userRepo *repositories.UserRepository, roleRepo *repositories.RoleRepository) *UserService {
	return &UserService{UserRepo: userRepo, RoleRepo: roleRepo}
}

func (s *UserService) Register(username, email, password string, roleId uint) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{Username: username, Email: email, Password: string(hashedPassword)}

	role, err := s.RoleRepo.FindById(roleId)
	if err != nil {
		return models.User{}, err
	}
	user.Role = *role

	if err := s.UserRepo.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Compare stored and provided passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	userRole, err := s.UserRepo.GetUserWithRole(user.ID)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Generate JWT with user roles
	token, err := utils.GenerateJWT(userRole.ID, userRole.Role.Name)
	if err != nil {
		return "", err
	}

	return user.Role.Name + token, nil
}

func (s *UserService) GetCurrentUser(id string) (models.User, error) {
	user, err := s.UserRepo.FindById(id)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	return user, err
}
