package handlers

import (
	"net/http"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/services"
	"strings"

	"github.com/gin-gonic/gin"
)

type ForgotPasswordHandler struct {
	forgotPasswordService *services.ForgotPasswordService
}

func NewForgotPasswordHandler(forgotPasswordService *services.ForgotPasswordService) *ForgotPasswordHandler {
	return &ForgotPasswordHandler{
		forgotPasswordService: forgotPasswordService,
	}
}

// RequestForgotPassword godoc
// @Summary      Request password reset OTP
// @Description  Sends or records an OTP for the given email if the user exists.
// @Tags         forgot-password
// @Accept       json
// @Produce      json
// @Param        body  body      dto.RequestForgotPasswordDTO  true  "Registered email"
// @Success      200   {object}  dto.Response
// @Failure      400   {object}  dto.Response
// @Failure      404   {object}  dto.Response
// @Failure      500   {object}  dto.Response
// @Router       /forgot-password/request [post]
func (u ForgotPasswordHandler) RequestForgotPassword(ctx *gin.Context) {
	var body dto.RequestForgotPasswordDTO
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	result, err := u.forgotPasswordService.RequestForgotPassword(body.Email)
	if err != nil {
		msg := err.Error()
		status := http.StatusInternalServerError
		switch {
		case strings.Contains(msg, "Invalid email format"):
			status = http.StatusBadRequest
		case strings.Contains(msg, "Email not found"):
			status = http.StatusNotFound
		}
		ctx.JSON(status, dto.Response{
			Success: false,
			Message: msg,
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Request forgot password success !",
		Data:    result,
	})
}

// ResetPassword godoc
// @Summary      Reset password with OTP
// @Description  Validates OTP and sets a new password for the account.
// @Tags         forgot-password
// @Accept       json
// @Produce      json
// @Param        body  body      dto.ResetForgotPasswordDTO  true  "Email, OTP, and new password"
// @Success      200   {object}  dto.Response
// @Failure      400   {object}  dto.Response
// @Failure      404   {object}  dto.Response
// @Failure      500   {object}  dto.Response
// @Router       /forgot-password/reset [post]
func (u ForgotPasswordHandler) ResetPassword(ctx *gin.Context) {
	var body dto.ResetForgotPasswordDTO
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	err = u.forgotPasswordService.ResetPassword(body)
	if err != nil {
		msg := err.Error()
		status := http.StatusInternalServerError
		switch {
		case strings.Contains(msg, "Invalid email format"),
			strings.Contains(msg, "Password too weak"),
			strings.Contains(msg, "missmatch"),
			strings.Contains(msg, "Failed to trim otp"),
			strings.Contains(msg, "Failed to convert otp"),
			strings.Contains(msg, "OTP is invalid"):
			status = http.StatusBadRequest
		case strings.Contains(msg, "Email not found"),
			strings.Contains(msg, "OTP not found"):
			status = http.StatusNotFound
		}
		ctx.JSON(status, dto.Response{
			Success: false,
			Message: msg,
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Reset password success !",
		Data:    nil,
	})
}
