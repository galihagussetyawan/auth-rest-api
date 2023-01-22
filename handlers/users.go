package handlers

import (
	"auth-rest-api/configs"
	"auth-rest-api/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CreateUser(c *fiber.Ctx) error {
	userCollection := configs.ConnectDB().Database("go-db").Collection("users")
	userReqBody := models.User{}

	err := c.BodyParser(&userReqBody)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status_code": fiber.StatusBadRequest,
			"message":     err.Error(),
		})
	}

	hash, err := hashPassword(userReqBody.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status_code": fiber.StatusBadRequest,
			"message":     err.Error(),
		})
	}

	newUser := models.User{
		ID:        primitive.NewObjectID(),
		Firstname: userReqBody.Firstname,
		Lastname:  userReqBody.Lastname,
		Email:     userReqBody.Email,
		Password:  hash,
	}

	result, err := userCollection.InsertOne(c.Context(), &newUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status_code": fiber.StatusBadRequest,
			"message":     err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status_code": fiber.StatusCreated,
		"message":     "success to create an account",
		"data":        result,
	})
}
