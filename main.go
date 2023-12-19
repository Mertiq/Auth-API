package main

import (
	"Auth-API/infrastracture"
	"Auth-API/infrastracture/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	db := database.NewPostgresConn()

	app := fiber.New()
	app.Use(cors.New())

	infrastracture.SetupRoutes(db, app)
	app.Listen(":3000")

}
