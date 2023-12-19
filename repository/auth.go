package repository

import (
	"Auth-API/domain"
	"Auth-API/infrastracture/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthRepository struct {
	Database *gorm.DB
}

func (repository *AuthRepository) Register(ctx *fiber.Ctx, user *domain.User) error {
	if err := repository.Database.Create(&user).Error; err != nil {
		return errors.NewInternalServerError(ctx, err)
	}
	return nil
}
