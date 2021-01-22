package main

import (
	"timeo-api/config"
	"timeo-api/database"
	"timeo-api/router"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	database.ConnectDB()

	router.SetupRoutes(app)
	panic(app.Listen(config.Config("PORT")))
}
