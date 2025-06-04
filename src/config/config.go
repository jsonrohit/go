package config

import (
	"fiber/src/utils"

	"github.com/gofiber/fiber/v2"
)

// Config holds application configuration
type Config struct {
	Port string
	Env  string
}

// New creates a new configuration
func New() *Config {
	return &Config{
		Port: "3000",
		Env:  "development",
	}
}

// ErrorHandler handles application errors
func ErrorHandler(c *fiber.Ctx, err error) error {
	// Default 500 statuscode
	code := fiber.StatusInternalServerError

	// Check if it's a fiber.Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(utils.ErrorResponse(
		"Internal Server Error",
		err.Error(),
	))
}