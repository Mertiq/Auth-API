package response

import "github.com/golang-jwt/jwt"

type JWTClaim struct {
	Sub uint `json:"sub"`
	jwt.StandardClaims
}
