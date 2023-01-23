package routes

import (
	"auth-rest-api/handlers"
	"auth-rest-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(a *fiber.App) {
	route := a.Group("/api/v1/auth")

	route.Post("/login", handlers.Login)
	route.Post("/refresh-token", middleware.LoginRequired, handlers.RefreshToken)
}
