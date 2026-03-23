package dto

import "time"

type UserResponseDTO struct {
	Id          int       `json:"id"`
	User_avatar string    `json:"user_avatar"`
	First_name  string    `json:"first_name"`
	Last_name   string    `json:"last_name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
	Verified    bool      `json:"verified"`
	Role_name   string    `json:"role_name"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

type CreateUserDTO struct {
	First_name       string `json:"first_name"`
	Last_name        string `json:"last_name"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Address          string `json:"address"`
	Password         string `json:"password"`
	Password_confirm string `json:"password_confirm"`
}

type UpdateUserProfileDTO struct {
	Id          int    `json:"id"`
	First_name  string `json:"first_name"`
	Last_name   string `json:"last_name"`
	Address     string `json:"address"`
	User_avatar string `json:"user_avatar"`
}
