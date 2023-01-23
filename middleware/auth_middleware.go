package middleware

import (
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func getBearerToken(c *fiber.Ctx) string {
	tokenString := c.Get("Authorization")
	tokenSplit := strings.Split(tokenString, "Bearer")
	if len(tokenSplit) != 2 {
		return ""
	}
	tokenString = strings.TrimSpace(tokenSplit[1])

	return tokenString
}

func LoginRequired(c *fiber.Ctx) error {
	tokenString := getBearerToken(c)
	if tokenString == "" {
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status_code": fiber.StatusUnauthorized,
				"message":     "require bearer token",
			})
		}
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status_code": fiber.StatusUnauthorized,
			"message":     err.Error(),
		})
	}

	claims := token.Claims.(jwt.MapClaims)
	userId := claims["id"]

	c.Locals("user_id", userId)
	return c.Next()
}

func RequiredAdmin(c *fiber.Ctx) error {
	log.Println("asdasd")
	return c.Next()
}
