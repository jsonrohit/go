// internal/repository/user_repository.go
package repository

import (
	"errors"
	"fiber/src/models"
	"time"
)

// UserRepository defines the interface for user data operations
type UserRepository interface {
	GetAll() ([]*models.User, error)
	GetByID(id int) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Create(user *models.User) (*models.User, error)
	Update(id int, user *models.User) (*models.User, error)
	Delete(id int) error
	EmailExists(email string, excludeID int) (bool, error)
}

// InMemoryUserRepository implements UserRepository using in-memory storage
type InMemoryUserRepository struct {
	users  []*models.User
	nextID int
}

// NewInMemoryUserRepository creates a new in-memory user repository
func NewInMemoryUserRepository() UserRepository {
	repo := &InMemoryUserRepository{
		users:  make([]*models.User, 0),
		nextID: 1,
	}

	// Initialize with sample data
	repo.initSampleData()

	return repo
}

// initSampleData adds some initial users
func (r *InMemoryUserRepository) initSampleData() {
	now := time.Now()

	sampleUsers := []*models.User{
		{
			ID:        1,
			Name:      "John Doe",
			Email:     "john@example.com",
			Age:       30,
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			ID:        2,
			Name:      "Jane Smith",
			Email:     "jane@example.com",
			Age:       25,
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			ID:        3,
			Name:      "Bob Johnson",
			Email:     "bob@example.com",
			Age:       35,
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	r.users = sampleUsers
	r.nextID = 4
}

// GetAll returns all users
func (r *InMemoryUserRepository) GetAll() ([]*models.User, error) {
	return r.users, nil
}

// GetByID returns a user by ID
func (r *InMemoryUserRepository) GetByID(id int) (*models.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

// GetByEmail returns a user by email
func (r *InMemoryUserRepository) GetByEmail(email string) (*models.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

// Create creates a new user
func (r *InMemoryUserRepository) Create(user *models.User) (*models.User, error) {
	user.ID = r.nextID
	r.nextID++

	r.users = append(r.users, user)
	return user, nil
}

// Update updates an existing user
func (r *InMemoryUserRepository) Update(id int, updatedUser *models.User) (*models.User, error) {
	for i, user := range r.users {
		if user.ID == id {
			updatedUser.ID = id
			updatedUser.CreatedAt = user.CreatedAt // Preserve creation time
			updatedUser.UpdatedAt = time.Now()

			r.users[i] = updatedUser
			return updatedUser, nil
		}
	}
	return nil, errors.New("user not found")
}

// Delete deletes a user by ID
func (r *InMemoryUserRepository) Delete(id int) error {
	for i, user := range r.users {
		if user.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

// EmailExists checks if an email already exists (excluding a specific user ID)
func (r *InMemoryUserRepository) EmailExists(email string, excludeID int) (bool, error) {
	for _, user := range r.users {
		if user.Email == email && user.ID != excludeID {
			return true, nil
		}
	}
	return false, nil
}
