package main

import (
	"going/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	//HOME
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	config.InitMongo()
	app.Listen(":3000")
}
