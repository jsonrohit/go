// internal/models/user.go
package models

import "time"

// User represents a user in the system
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required,min=2,max=50"`
	Email     string    `json:"email" validate:"required,email"`
	Age       int       `json:"age" validate:"min=1,max=120"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUserRequest represents the request payload for creating a user
type CreateUserRequest struct {
	Name  string `json:"name" validate:"required,min=2,max=50"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"min=1,max=120"`
}

// UpdateUserRequest represents the request payload for updating a user
type UpdateUserRequest struct {
	Name  string `json:"name" validate:"required,min=2,max=50"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"min=1,max=120"`
}

// ToUser converts CreateUserRequest to User
func (r *CreateUserRequest) ToUser() *User {
	now := time.Now()
	return &User{
		Name:      r.Name,
		Email:     r.Email,
		Age:       r.Age,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// ToUser converts UpdateUserRequest to User
func (r *UpdateUserRequest) ToUser() *User {
	return &User{
		Name:      r.Name,
		Email:     r.Email,
		Age:       r.Age,
		UpdatedAt: time.Now(),
	}
}