package models

import "time"

type Cart struct {
	Id        int        `json:"id" db:"id"`
	UserId    int        `json:"user_id" db:"user_id"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}