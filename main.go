package main

import (
	"auth-rest-api/configs"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	app.Listen(":8080")
}
