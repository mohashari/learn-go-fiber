package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	// Basic Routing
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	//Parameter
	app.Get("/:name", func(c *fiber.Ctx) error {
		return c.SendString("name: " + c.Params("name"))
	})

	//Optional Parameter
	app.Get("/nama/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hello: " + c.Params("name"))
		}
		return c.SendString("No name found")
	})

	app.Listen(":3000")
}
