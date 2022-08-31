package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ochoad24/quotes-go/connection"
	"github.com/ochoad24/quotes-go/models"
)

func GetQuoteController(c *fiber.Ctx) error {
	var quotes []models.Quote
	connection.DB.Where("is_del = ?", 0).Find(&quotes)

	return c.Status(fiber.StatusOK).JSON(&quotes)
}

func NewQuoteController(c *fiber.Ctx) error {
	var quote models.Quote
	if err := c.BodyParser(&quote); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	createdTask := connection.DB.Create(&quote)
	err := createdTask.Error

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).SendString("Cita creada")
}

func UpdateQuoteController(c *fiber.Ctx) error {
	var quote models.Quote
	id := c.Params("id")
	if err := c.BodyParser(&quote); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	connection.DB.Where("id = ?", id).Updates(&quote)

	return c.Status(fiber.StatusOK).SendString("Cita actualizada")
}

func DeleteQuoteController(c *fiber.Ctx) error {
	id := c.Params("id")
	var quote models.Quote

	found := connection.DB.Find(&quote, id)

	if found.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(found.Error.Error())
	}

	connection.DB.Model(&quote).Where("id = ?", id).Update("is_del", "0")

	return c.Status(fiber.StatusOK).SendString("Cita eliminada")

}
