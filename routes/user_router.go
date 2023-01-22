package routes

import (
	"auth-rest-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("/user", handlers.CreateUser)
}
