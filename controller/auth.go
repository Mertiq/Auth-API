package controller

import (
	"Auth-API/entity/dto/request"
	"Auth-API/infrastracture/errors"
	"Auth-API/repository"
	"Auth-API/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthController struct {
	AuthService *service.AuthService
}

func NewAuthController(database *gorm.DB) *AuthController {
	return &AuthController{
		AuthService: &service.AuthService{
			AuthRepository: &repository.AuthRepository{
				Database: database,
			},
		},
	}
}

func (controller *AuthController) Register(c *fiber.Ctx) error {
	var req request.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return errors.NewBadRequestError(c, err)
	}

	newUser, err := controller.AuthService.Register(c, req)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(newUser)
}

func (controller *AuthController) Login(c *fiber.Ctx) error {
	var req request.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return errors.NewBadRequestError(c, err)
	}

	token, err := controller.AuthService.Login(c, req)

	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(token)
}
