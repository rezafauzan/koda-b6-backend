package models

import "time"

type User struct {
	Id         int `json:"id"`
	Role_id    int `json:"role_id"`
	Verified   bool `json:"verified"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
