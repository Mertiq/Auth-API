package request

import (
	"Auth-API/entity/enums"
)

type RegisterRequest struct {
	UserName string         `json:"user_name"`
	Password string         `json:"password"`
	Address  string         `json:"address"`
	Role     enums.UserRole `json:"role"`
}

type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
