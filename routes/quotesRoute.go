package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ochoad24/quotes-go/controllers"
)

func QuoteRoute(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		return controllers.GetQuoteController(c)
	})

	router.Post("/", func(c *fiber.Ctx) error {
		return controllers.NewQuoteController(c)
	})

	router.Put("/:id", func(c *fiber.Ctx) error {
		return controllers.UpdateQuoteController(c)
	})

	router.Delete("/:id", func(c *fiber.Ctx) error {
		return controllers.DeleteQuoteController(c)
	})
}
