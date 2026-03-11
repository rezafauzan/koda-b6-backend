package models

import "time"

type Forgot_Password struct {
	id         int
	email      string
	code_otp   int
	expired_at time.Time
	created_at time.Time
	updated_at time.Time
}
