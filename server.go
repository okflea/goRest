package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	//HOME
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// WILDCARD
	// GET http://localhost:3000/api/user/john
	app.Get("/api/*", func(c *fiber.Ctx) error {
		return c.SendString("API path: " + c.Params("*"))
		// => API path: user/john
	})

	//PARAMS
	// GET http://localhost:8080/hello%20world
	app.Get("/:value", func(c *fiber.Ctx) error {
		return c.SendString("value: " + c.Params("value"))
		// => Get request with value: hello world
	})

	app.Listen(":3000")
}
