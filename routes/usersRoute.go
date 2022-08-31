package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ochoad24/quotes-go/controllers"
)

func UserRoute(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		return controllers.GetUserController(c)
	})

	router.Post("/", func(c *fiber.Ctx) error {
		return controllers.NewUserController(c)
	})
}
