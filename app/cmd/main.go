package main

import (
	"ecommerce/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()

	app.Get(
		"/", func(c *fiber.Ctx) error {
			return c.SendString("Hello, how are you?")
		},
	)

	app.Listen(":3000")
}
