package handler

import (
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": "LOL"})
}
