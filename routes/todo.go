package routes

import (
	"github.com/brenobaptista/library-go-fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func TodoRoute(route fiber.Router) {
	route.Get("", controllers.GetTodos)
}
