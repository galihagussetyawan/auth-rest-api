package routes

import (
	"auth-rest-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(a *fiber.App) {
	route := a.Group("/api/v1/auth")

	route.Get("/login", handlers.Login)
}
