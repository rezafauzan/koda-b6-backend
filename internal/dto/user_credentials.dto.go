package dto

import (
	"rezafauzan/koda-b6-golang/internal/models"
	"time"
)

type UserCredentialResponseDTO struct {
	Id         int       `json:"id"`
	User_id    int       `json:"user_id"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type CreateUserCredentialDTO struct {
	User_id  int    `json:"user_id"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UpdateUserCredentialDTO struct {
	Id       int    `json:"id"`
	User_id  int    `json:"user_id"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func UserCredentialResponseFromModel(m models.UserCredential) UserCredentialResponseDTO {
	return UserCredentialResponseDTO{
		Id: m.Id, User_id: m.UserId, Email: m.Email, Phone: m.Phone, Password: m.Password,
		Created_at: m.CreatedAt, Updated_at: m.UpdatedAt,
	}
}