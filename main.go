package main

import (
	"auth-rest-api/configs"
	"auth-rest-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	configs.ConnectDB()

	//routers
	routes.UserRoutes(app)

	app.Listen(":8080")
}
