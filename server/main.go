package server

import (
	handlers "todo/api"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(server *fiber.App) {
	server.Get("/", handlers.GetHomePage)
	server.Post("/todo", handlers.CreateTodo)
	server.Get("/todos", handlers.GetTodos)
}

func InitServer() {
	app := fiber.New()

	SetRoutes(app)

	app.Listen(":3030")
}
