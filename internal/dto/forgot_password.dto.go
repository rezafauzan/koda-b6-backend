package dto

type ForgotPasswordRequest struct{
	Email string
}

type ForgotPasswordUpdateRequest struct{
	Email string
	code_otp int
	NewPassword string
	ConfirmNewPassword string
}