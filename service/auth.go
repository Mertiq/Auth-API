package service

import (
	"Auth-API/domain"
	"Auth-API/domain/dto/request"
	"Auth-API/repository"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct {
	AuthRepository *repository.AuthRepository
}

func (service *AuthService) Register(ctx *fiber.Ctx, request request.RegisterRequest) (domain.User, error) {
	password, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	newUser := domain.User{
		UserName:  request.UserName,
		Password:  password,
		Address:   request.Address,
		Role:      request.Role,
		CreatedAt: time.Now(),
	}

	err := service.AuthRepository.Register(ctx, &newUser)

	return newUser, err
}
