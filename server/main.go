package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", getGreeting)

	log.Fatal(app.Listen(":3000"))
}
