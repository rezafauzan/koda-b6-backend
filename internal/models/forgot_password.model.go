package models

import "time"

type ForgotPassword struct {
	Id         int `json:"id"`
	Email      string `json:"email"`
	Code_otp   int `json:"code_otp"`
	Expired_at time.Time `json:"expired_at"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
