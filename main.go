package main

import (
	"github.com/steffengeipel/timeo-api/config"
	"github.com/steffengeipel/timeo-api/router"

	"github.com/gofiber/fiber/v2"
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

	router.SetupRoutes(app)
	app.Listen(config.Config("PORT"))

}

func GetToDos(ctx *fiber.Ctx) {
	ctx.Status(fiber.StatusOK).JSON(todos)
}
