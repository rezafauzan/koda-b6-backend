package models

import "time"

type UserProfile struct {
	Id         int       `json:"id" db:"id"`
	UserId     int       `json:"user_id" db:"user_id"`
	UserAvatar string    `json:"user_avatar" db:"user_avatar"`
	FirstName  string    `json:"first_name" db:"first_name"`
	LastName   string    `json:"last_name" db:"last_name"`
	Address    string    `json:"address" db:"address"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
