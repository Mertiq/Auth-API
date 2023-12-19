package controller

import (
	"Auth-API/domain/dto/request"
	"Auth-API/infrastracture/errors"
	"Auth-API/service"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService service.AuthService
}

func (controller *AuthController) Register(c *fiber.Ctx) error {
	var req request.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return errors.NewBadRequestError(c, err)
	}

	newUser, err := controller.authService.Register(req)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(newUser)
}
