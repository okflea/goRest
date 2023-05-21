package main

import (
	"going/config"

	"github.com/gofiber/fiber/v2"
	// "goingNowhere/config"
)

// func initMongo() {
// 	// Setup the mgm default config
// 	errEnv := godotenv.Load()
// 	if errEnv != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	// fmt.Println("uri:", os.Getenv("MONGOURI"))
// 	errMongo := mgm.SetDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI(os.Getenv("MONGOURI")))
//
// 	if errMongo != nil {
// 		log.Fatal("Error connecting to mongo")
// 	} else {
// 		fmt.Println("Connected to db")
// 	}
// }

func main() {
	app := fiber.New()
	//HOME
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	config.InitMongo()
	// WILDCARD
	// GET http://localhost:3000/api/user/john
	// app.Get("/api/*", func(c *fiber.Ctx) error {
	// 	return c.SendString("API path: " + c.Params("*"))
	// 	// => API path: user/john
	// })
	//
	// //PARAMS
	// // GET http://localhost:8080/hello%20world
	// app.Get("/:value", func(c *fiber.Ctx) error {
	// 	return c.SendString("value: " + c.Params("value"))
	// 	// => Get request with value: hello world
	// })

	app.Listen(":3000")
}
