package infrastracture

import (
	"Auth-API/service"
	"github.com/gofiber/fiber/v2"
)

func ProtectWithJWT(handler fiber.Handler) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		headers := ctx.GetReqHeaders()
		headerArr, exists := headers["Authorization"]
		token := headerArr[0]

		if !exists {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		if err := service.ValidateToken(token); err != nil {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		userID := service.ExtractUserIDFromToken(token)
		if userID == 0 {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		ctx.Locals("userID", userID)
		return handler(ctx)
	}
}
