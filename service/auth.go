package service

import (
	"Auth-API/domain"
	"Auth-API/domain/dto/request"
	"Auth-API/repository"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct {
	authRepository repository.AuthRepository
}

func (service *AuthService) Register(request request.RegisterRequest) (domain.User, error) {
	password, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	newUser := domain.User{
		UserName:  request.UserName,
		Password:  password,
		Address:   request.Address,
		Role:      request.Role,
		CreatedAt: time.Now(),
	}

	err := service.authRepository.Register(&newUser)

	return newUser, err
}
