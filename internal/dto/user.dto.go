package dto

import "time"

type User struct {
	id          int
	user_avatar string
	first_name  string
	last_name   string
	email       string
	phone       string
	address     string
	verified    bool
	role_name   string
	created_at  time.Time
	updated_at  time.Time
}
