package dto

import (
	"rezafauzan/koda-b6-golang/internal/models"
	"time"
)

type UserProfileResponseDTO struct {
	Id          int       `json:"id"`
	User_id     int       `json:"user_id"`
	User_avatar string    `json:"user_avatar"`
	First_name  string    `json:"first_name"`
	Last_name   string    `json:"last_name"`
	Address     string    `json:"address"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

type CreateUserProfileDTO struct {
	User_id     int    `json:"user_id"`
	User_avatar string `json:"user_avatar"`
	First_name  string `json:"first_name"`
	Last_name   string `json:"last_name"`
	Address     string `json:"address"`
}

type UpdateUserProfileEntityDTO struct {
	Id          int    `json:"id"`
	User_id     int    `json:"user_id"`
	User_avatar string `json:"user_avatar"`
	First_name  string `json:"first_name"`
	Last_name   string `json:"last_name"`
	Address     string `json:"address"`
}

func UserProfileResponseFromModel(m models.UserProfile) UserProfileResponseDTO {
	return UserProfileResponseDTO{
		Id: m.Id, User_id: m.UserId, User_avatar: m.UserAvatar, First_name: m.FirstName, Last_name: m.LastName,
		Address: m.Address, Created_at: m.CreatedAt, Updated_at: m.UpdatedAt,
	}
}