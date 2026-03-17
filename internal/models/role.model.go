package models

import "time"

type Role struct {
	Id         int       `json:"id"`
	Role_name  string    `json:"role_name"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
