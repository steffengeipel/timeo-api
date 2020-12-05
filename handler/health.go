package handler

import (
	"github.com/gofiber/fiber/v2"
)

// GetAPIHealth Check the Health from API
func GetAPIHealth(c *fiber.Ctx) error {
	t := struct {
		test string
	}{
		test: "LOL",
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}
