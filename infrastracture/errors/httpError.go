package errors

import "github.com/gofiber/fiber/v2"

func NewBadRequestError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusBadRequest).JSON(err.Error())
}

func NewInternalServerError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
}
