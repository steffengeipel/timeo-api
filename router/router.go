package router

import (
	"timeo-api/handler"
	"timeo-api/router/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	// User
	routes.SetupUserRoutes(api)

	health := api.Group("/health")
	health.Get("/api", handler.GetAPIHealth)
}
