package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/steffengeipel/timeo-api/config"
	"github.com/steffengeipel/timeo-api/database"
	"github.com/steffengeipel/timeo-api/router"
)

type ToDo struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var todos = []ToDo{
	{Id: 1, Name: "Walk the Dog", Completed: false},
	{Id: 2, Name: "Walk the Cat", Completed: false},
}

func main() {

	app := fiber.New()

	database.ConnectDB()

	router.SetupRoutes(app)
	app.Listen(config.Config("PORT"))

	defer database.DB.Close()

}

func GetToDos(ctx *fiber.Ctx) {
	ctx.Status(fiber.StatusOK).JSON(todos)
}
