package request

import "Auth-API/domain/enums"

type RegisterRequest struct {
	UserName string         `json:"user_name"`
	Password string         `json:"password"`
	Address  string         `json:"address"`
	Role     enums.UserRole `json:"role"`
}
