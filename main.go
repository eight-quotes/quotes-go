package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ochoad24/quotes-go/connection"
	"github.com/ochoad24/quotes-go/routes"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	connection.DBConnection()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://magical-malabi-19c8d6.netlify.app",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("READY!")
	})

	routes.QuoteRoute(app.Group("/quotes"))
	routes.UserRoute(app.Group("/users"))

	app.Listen("0.0.0.0:" + port)
	fmt.Println("Server listening on " + port)
}
