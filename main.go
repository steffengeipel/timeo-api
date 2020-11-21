package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-todos/router"
)


type ToDo struct {
	Id	 				int	 			`json:"id"`
	Name 				string		`json:"name"`
	Completed		bool			`json:"completed"`
}

var todos = []ToDo{
	{Id: 1, Name: "Walk the Dog", Completed: false},
	{Id: 2, Name: "Walk the Cat", Completed: false},
}

func main() {
	
  app := fiber.New()

	router.SetupRoutes(app)
  app.Listen(":3000")
		
}



func GetToDos(ctx *fiber.Ctx){
	ctx.Status(fiber.StatusOK).JSON(todos)
}