package routes

import (
	"fiber/src/handlers"
	"fiber/src/repository"
	services "fiber/src/services/users"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// SetupRoutes configures all application routes
func SetupRoutes(app *fiber.App) {
	// Middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "${time} ${status} - ${method} ${path} - ${latency}\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
	}))

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "OK",
			"message": "Server is running",
		})
	})

	// Initialize dependencies
	userRepo := repository.NewInMemoryUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// API routes
	api := app.Group("/api/v1")

	// User routes
	userRoutes := api.Group("/users")
	userRoutes.Get("/", userHandler.GetUsers)         // GET /api/v1/users
	userRoutes.Get("/:id", userHandler.GetUser)       // GET /api/v1/users/:id
	userRoutes.Post("/", userHandler.CreateUser)      // POST /api/v1/users
	userRoutes.Put("/:id", userHandler.UpdateUser)    // PUT /api/v1/users/:id
	userRoutes.Delete("/:id", userHandler.DeleteUser) // DELETE /api/v1/users/:id

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Route not found",
		})
	})
}
