package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/steffengeipel/timeo-api/config"
	"github.com/steffengeipel/timeo-api/database"
	"github.com/steffengeipel/timeo-api/router"
)

func main() {

	app := fiber.New()

	database.ConnectDB()

	router.SetupRoutes(app)
	app.Listen(config.Config("PORT"))
}
