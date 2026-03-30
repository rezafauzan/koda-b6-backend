package dto

import (
	"time"
)

type UserCredentialResponseDTO struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserCredentialResponseWithoutPasswordDTO struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateUserCredentialDTO struct {
	Id              int    `json:"id"`
	UserId          int    `json:"user_id"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
