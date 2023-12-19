package entity

import (
	"Auth-API/entity/enums"
	"time"
)

type User struct {
	Id        uint           `json:"id"`
	UserName  string         `json:"user_name"`
	Password  []byte         `json:"password"`
	Address   string         `json:"address"`
	Role      enums.UserRole `json:"role"`
	CreatedAt time.Time      `json:"created_at"`
}
