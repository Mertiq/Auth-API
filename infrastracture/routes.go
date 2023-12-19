package infrastracture

import (
	"Auth-API/controller"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(database *gorm.DB, app *fiber.App) {
	authController := controller.NewAuthController(database)

	app.Post("/auth/register", authController.Register)
}
