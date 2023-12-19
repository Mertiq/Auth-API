package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	//ConnectDatabase()

	app := fiber.New()
	app.Use(cors.New())

	//SetupRoutes(app)
	app.Listen(":3000")

}
