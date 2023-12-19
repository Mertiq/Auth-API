package service

import (
	"Auth-API/entity/dto/response"
	"errors"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

func CreateAccessToken(id uint) (string, error) {
	claims := response.JWTClaim{
		Sub: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func ValidateToken(tokenString string) error {
	token, err := ParseTokenWithClaims(tokenString)

	if err != nil {
		return err
	}

	claims, ok := token.Claims.(*response.JWTClaim)
	if !ok {
		return errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return errors.New("token expired")
	}

	return nil
}

func ExtractUserIDFromToken(tokenString string) uint {
	token, err := ParseTokenWithClaims(tokenString)

	if err != nil {
		return 0
	}

	if claims, ok := token.Claims.(*response.JWTClaim); ok {
		return claims.Sub
	}

	return 0
}
func ParseTokenWithClaims(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &response.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
}
