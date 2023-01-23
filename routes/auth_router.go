package routes

import (
	"auth-rest-api/handlers"
	"auth-rest-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(a *fiber.App) {
	v1 := a.Group("/api/v1/auth")

	v1.Post("/login", handlers.Login)
	v1.Post("/refresh-token", middleware.LoginRequired, handlers.RefreshToken)
}
