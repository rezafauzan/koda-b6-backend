package dto

import (
	"time"
)

type UserProfileResponseDTO struct {
	Id          int       `json:"id"`
	UserId     int       `json:"user_id"`
	UserAvatar string    `json:"user_avatar"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Address     string    `json:"address"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateUserProfileDTO struct {
	User_id     int    `json:"user_id"`
	User_avatar string `json:"user_avatar"`
	First_name  string `json:"first_name"`
	Last_name   string `json:"last_name"`
	Address     string `json:"address"`
}

type UpdateUserProfileDTO struct {
	Id          int    `json:"id"`
	User_avatar string `json:"user_avatar"`
	First_name  string `json:"first_name"`
	Last_name   string `json:"last_name"`
	Address     string `json:"address"`
}