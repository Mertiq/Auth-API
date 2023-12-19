package repository

import (
	"Auth-API/domain"
	"Auth-API/infrastracture/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthRepository struct {
	database *gorm.DB
	ctx      *fiber.Ctx
}

func (repository *AuthRepository) Register(user *domain.User) error {
	if err := repository.database.Create(&user).Error; err != nil {
		return errors.NewInternalServerError(repository.ctx, err)
	}
	return nil
}
