package models

import "time"

type User struct {
	Id         int
	Role_id    int
	Verified   bool
	Created_at time.Time
	Updated_at time.Time
}
