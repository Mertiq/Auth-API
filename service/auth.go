package service

import (
	"Auth-API/entity"
	"Auth-API/entity/dto/request"
	"Auth-API/infrastracture/errors"
	"Auth-API/repository"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct {
	AuthRepository *repository.AuthRepository
}

func (service *AuthService) Register(ctx *fiber.Ctx, request request.RegisterRequest) (entity.User, error) {
	password, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	newUser := entity.User{
		UserName:  request.UserName,
		Password:  password,
		Address:   request.Address,
		Role:      request.Role,
		CreatedAt: time.Now(),
	}

	err := service.AuthRepository.CreateUser(ctx, &newUser)

	return newUser, err
}

func (service *AuthService) Login(c *fiber.Ctx, req request.LoginRequest) (string, error) {
	var token string
	user, err := service.AuthRepository.FindUserByCredential(c, req)
	if err != nil {
		return token, err
	}
	if user.Id == 0 {
		return token, c.Status(fiber.StatusNotFound).JSON("User Not Found")
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(req.Password)); err != nil {
		return token, errors.NewBadRequestError(c, err)
	}

	token, err = CreateAccessToken(user.Id)

	if err != nil {
		return token, errors.NewInternalServerError(c, err)
	}

	return token, err
}
