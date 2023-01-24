package routes

import (
	"auth-rest-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(a *fiber.App) {
	v1 := a.Group("/api/v1")

	v1.Post("/user", handlers.CreateUser)
}
