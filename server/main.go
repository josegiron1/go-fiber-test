package server

import (
	handlers "todo/api"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(server *fiber.App) {
	server.Get("/", handlers.GetGreeting)
}

func InitServer() {
	app := fiber.New()

	SetRoutes(app)

	app.Listen(":3030")
}
