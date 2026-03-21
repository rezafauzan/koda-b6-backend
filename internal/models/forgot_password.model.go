package models

import "time"

type ForgotPassword struct {
	Id        int       `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	CodeOtp   int       `json:"code_otp" db:"code_otp"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
