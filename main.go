package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ochoad24/quotes-go/connection"
	"github.com/ochoad24/quotes-go/routes"
)

func main() {

	connection.DBConnection()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("READY!")
	})

	routes.QuoteRoute(app.Group("/quotes"))
	routes.UserRoute(app.Group("/users"))

	app.Listen(":3000")
	fmt.Println("Server listening on 3000")
}
