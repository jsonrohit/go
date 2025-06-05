package services

import (
	"errors"
	"fiber/src/models"
	"fiber/src/repository"
	"fiber/src/utils"
)

// UserService defines the interface for user business logic
type UserService interface {
	GetAllUsers() ([]*models.User, error)
	GetUserByID(id int) (*models.User, error)
	CreateUser(req *models.CreateUserRequest) (*models.User, error)
	UpdateUser(id int, req *models.UpdateUserRequest) (*models.User, error)
	DeleteUser(id int) (*models.User, error)
}

// UserServiceImpl implements UserService
type UserServiceImpl struct {
	repo repository.UserRepository
}

// NewUserService creates a new user service
func NewUserService(repo repository.UserRepository) UserService {
	return &UserServiceImpl{
		repo: repo,
	}
}

// GetAllUsers retrieves all users
func (s *UserServiceImpl) GetAllUsers() ([]*models.User, error) {
	return s.repo.GetAll()
}

// GetUserByID retrieves a user by ID
func (s *UserServiceImpl) GetUserByID(id int) (*models.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}

	return s.repo.GetByID(id)
}

// CreateUser creates a new user
func (s *UserServiceImpl) CreateUser(req *models.CreateUserRequest) (*models.User, error) {
	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return nil, err
	}

	// Check if email already exists
	exists, err := s.repo.EmailExists(req.Email, 0)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("email already exists")
	}

	// Convert request to user model
	user := req.ToUser()

	// Create user
	return s.repo.Create(user)
}

// UpdateUser updates an existing user
func (s *UserServiceImpl) UpdateUser(id int, req *models.UpdateUserRequest) (*models.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return nil, err
	}

	// Check if user exists
	existingUser, err := s.repo.GetByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Check if email already exists (excluding current user)
	exists, err := s.repo.EmailExists(req.Email, id)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("email already exists")
	}

	// Convert request to user model
	updatedUser := req.ToUser()
	updatedUser.CreatedAt = existingUser.CreatedAt // Preserve creation time

	// Update user
	return s.repo.Update(id, updatedUser)
}

// DeleteUser deletes a user by ID
func (s *UserServiceImpl) DeleteUser(id int) (*models.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}

	// Get user before deleting (to return in response)
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Delete user
	if err := s.repo.Delete(id); err != nil {
		return nil, err
	}

	return user, nil
}
