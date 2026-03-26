package dto

import (
	"time"
)

type RoleResponseDTO struct {
	Id         int       `json:"id"`
	Role_name  string    `json:"role_name"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type CreateRoleDTO struct {
	Role_name string `json:"role_name"`
}

type UpdateRoleDTO struct {
	Id        int    `json:"id"`
	Role_name string `json:"role_name"`
}