// internal/handlers/user_handler.go
package handlers

import (
	"strconv"

	"fiber/src/models"
	"fiber/src/services/users"
	"fiber/src/utils"

	"github.com/gofiber/fiber/v2"
)

// UserHandler handles HTTP requests for users
type UserHandler struct {
	userService services.UserService
}

// NewUserHandler creates a new user handler
func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetUsers handles GET /users - Get all users
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			utils.ErrorResponse("Failed to retrieve users", err.Error()),
		)
	}

	return c.JSON(utils.SuccessResponse(users, "Users retrieved successfully"))
}

// GetUser handles GET /users/:id - Get user by ID
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			utils.ErrorResponse("Invalid user ID", "ID must be a number"),
		)
	}

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(
			utils.ErrorResponse("User not found", err.Error()),
		)
	}

	return c.JSON(utils.SuccessResponse(user, "User found"))
}

// CreateUser handles POST /users - Create new user
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			utils.ErrorResponse("Invalid request body", err.Error()),
		)
	}

	// Create user
	user, err := h.userService.CreateUser(&req)
	if err != nil {
		statusCode := fiber.StatusInternalServerError
		if err.Error() == "email already exists" {
			statusCode = fiber.StatusConflict
		} else if err.Error() == "validation failed" {
			statusCode = fiber.StatusBadRequest
		}

		return c.Status(statusCode).JSON(
			utils.ErrorResponse("Failed to create user", err.Error()),
		)
	}

	return c.Status(fiber.StatusCreated).JSON(
		utils.SuccessResponse(user, "User created successfully"),
	)
}

// UpdateUser handles PUT /users/:id - Update user
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			utils.ErrorResponse("Invalid user ID", "ID must be a number"),
		)
	}

	var req models.UpdateUserRequest

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			utils.ErrorResponse("Invalid request body", err.Error()),
		)
	}

	// Update user
	user, err := h.userService.UpdateUser(id, &req)
	if err != nil {
		statusCode := fiber.StatusInternalServerError
		if err.Error() == "user not found" {
			statusCode = fiber.StatusNotFound
		} else if err.Error() == "email already exists" {
			statusCode = fiber.StatusConflict
		} else if err.Error() == "validation failed" || err.Error() == "invalid user ID" {
			statusCode = fiber.StatusBadRequest
		}

		return c.Status(statusCode).JSON(
			utils.ErrorResponse("Failed to update user", err.Error()),
		)
	}

	return c.JSON(utils.SuccessResponse(user, "User updated successfully"))
}

// DeleteUser handles DELETE /users/:id - Delete user
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			utils.ErrorResponse("Invalid user ID", "ID must be a number"),
		)
	}

	user, err := h.userService.DeleteUser(id)
	if err != nil {
		statusCode := fiber.StatusInternalServerError
		if err.Error() == "user not found" || err.Error() == "invalid user ID" {
			statusCode = fiber.StatusNotFound
		}

		return c.Status(statusCode).JSON(
			utils.ErrorResponse("Failed to delete user", err.Error()),
		)
	}

	return c.JSON(utils.SuccessResponse(user, "User deleted successfully"))
}