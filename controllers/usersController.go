package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ochoad24/quotes-go/connection"
	"github.com/ochoad24/quotes-go/models"
	"golang.org/x/crypto/bcrypt"
)

func GetUserController(c *fiber.Ctx) error {
	var user []models.User
	connection.DB.Find(&user)
	return c.Status(fiber.StatusOK).JSON(&user)
}

func NewUserController(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	password, err := hashPassword(user.Password)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	user.Password = password

	createdTask := connection.DB.Create(&user)

	if createdTask.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(createdTask.Error.Error())
	}

	return c.Status(fiber.StatusOK).SendString("Usuario creado")
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
