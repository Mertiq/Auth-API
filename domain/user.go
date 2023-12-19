package domain

import "time"

type UserRole string

const (
	Seller   UserRole = "Seller"
	Customer UserRole = "Customer"
)

type User struct {
	Id        uint      `json:"id"`
	UserName  string    `json:"user_name"`
	Password  []byte    `json:"password"`
	Address   string    `json:"address"`
	Role      UserRole  `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
