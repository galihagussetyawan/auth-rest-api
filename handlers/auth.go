package handlers

import (
	"auth-rest-api/models"
	"auth-rest-api/repository"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func createToken(user models.User) string {
	claims := jwt.MapClaims{}
	claims["email"] = user.Email
	claims["id"] = user.ID
	claims["expires"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt, _ := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	return jwt
}

func createRefreshToken(user models.User) string {
	claims := jwt.MapClaims{}
	claims["email"] = user.Email
	claims["id"] = user.ID
	claims["expires"] = time.Now().Add(time.Hour * 800).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt, _ := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	return jwt
}

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}
	userRepo := repository.NewUserRepository()

	inputLogin := &LoginInput{}
	err := c.BodyParser(&inputLogin)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status_code": fiber.StatusBadRequest,
			"message":     err.Error(),
		})
	}

	userData, err := userRepo.GetUserByEmail(c.Context(), inputLogin.Identity)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status_code": fiber.StatusNotFound,
			"message":     "user not found",
		})
	}

	if !CheckPasswordHash(inputLogin.Password, userData.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status_code": fiber.StatusUnauthorized,
			"message":     "invalid password",
		})
	}

	accessToken := createToken(*userData)
	refreshToken := createRefreshToken(*userData)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status_code": fiber.StatusOK,
		"data": fiber.Map{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
			"token_type":    "bearer",
		},
	})
}
