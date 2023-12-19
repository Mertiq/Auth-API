package repository

import (
	"Auth-API/entity"
	"Auth-API/entity/dto/request"
	"Auth-API/infrastracture/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthRepository struct {
	Database *gorm.DB
}

func (repository *AuthRepository) CreateUser(ctx *fiber.Ctx, user *entity.User) error {
	if err := repository.Database.Create(&user).Error; err != nil {
		return errors.NewInternalServerError(ctx, err)
	}
	return nil
}

func (repository *AuthRepository) FindUserByCredential(ctx *fiber.Ctx, credential request.LoginRequest) (*entity.User, error) {
	var user *entity.User
	if err := repository.Database.Where("user_name = ?", credential.UserName).First(&user).Error; err != nil {
		return user, errors.NewInternalServerError(ctx, err)
	}
	return user, nil
}
