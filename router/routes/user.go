package routes

import (
	"timeo-api/handler"
	"timeo-api/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupUserRoutes loads the routs for the User
func SetupUserRoutes(api fiber.Router) {

	user := api.Group("/user")
	user.Get("/me", middleware.Protected(), handler.GetMyUserData)
	user.Post("/", handler.CreateUser)
	// user.Patch("/:id", middleware.Protected(), handler.UpdateUser)
	// user.Delete("/:id", middleware.Protected(), handler.DeleteUser)

}
