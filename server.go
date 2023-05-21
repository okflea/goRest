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
	app.Get("/create", func(c *fiber.Ctx) error {
		config.CreateUser("Rohit")
		return c.SendString("create!")
	})

	app.Get("/find", func(c *fiber.Ctx) error {
		config.FindUser()
		return c.SendString("find!")
	})
	
	config.InitMongo()
	app.Listen(":3000")
}
