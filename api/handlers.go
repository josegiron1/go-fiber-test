package handlers

import (
	"fmt"
	"todo/api/models"
	"todo/api/services"
	"todo/templates/components"
	"todo/templates/pages"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

var todoService = services.NewInMemoryTodoService()

func GetHomePage(c *fiber.Ctx) error {
	return adaptor.HTTPHandler(templ.Handler(pages.Page()))(c)
}

func CreateTodo(c *fiber.Ctx) error {
	fmt.Println(c)
	newTodo := new(models.Todo)

	if err := c.BodyParser(newTodo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON.",
		})
	}

	createdTodo, err := todoService.CreateTodo(newTodo)
	fmt.Println(createdTodo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create Todo",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(createdTodo)
}

func GetTodos(c *fiber.Ctx) error {
	todos, err := todoService.GetTodos()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get Todos",
		})
	}
	return adaptor.HTTPHandler(templ.Handler(components.TodosList(todos)))(c)
}
