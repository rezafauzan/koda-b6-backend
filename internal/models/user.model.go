package models

import "time"

type User struct {
	Id        int       `json:"id" db:"id"`
	RoleId    int       `json:"role_id" db:"role_id"`
	Verified  bool      `json:"verified" db:"verified"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
