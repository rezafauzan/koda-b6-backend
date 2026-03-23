package dto

type RequestForgotPasswordDTO struct {
	Email string `json:"email"`
}

type ResetForgotPasswordDTO struct {
	Email            string `json:"email"`
	Otp              string `json:"otp"`
	Code_otp         string `json:"code_otp"`
	New_password     string `json:"new_password"`
	Password_confirm string `json:"password_confirm"`
}
