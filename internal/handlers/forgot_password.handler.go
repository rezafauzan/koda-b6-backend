package handlers

import (
	"net/http"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/services"

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

func (u ForgotPasswordHandler) RequestForgotPassword(ctx *gin.Context) {
	var body dto.RequestForgotPasswordDTO
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success:  false,
			Messages: err.Error(),
			Results:  nil,
		})
		return
	}

	result, err := u.forgotPasswordService.RequestForgotPassword(body.Email)
	if err != nil {
		ctx.JSON(http.StatusOK, dto.Response{
			Success:  false,
			Messages: err.Error(),
			Results:  nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success:  true,
		Messages: "Request forgot password success !",
		Results:  result,
	})
}

func (u ForgotPasswordHandler) ResetPassword(ctx *gin.Context) {
	var body dto.ResetForgotPasswordDTO
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success:  false,
			Messages: err.Error(),
			Results:  nil,
		})
		return
	}

	err = u.forgotPasswordService.ResetPassword(body)
	if err != nil {
		ctx.JSON(http.StatusOK, dto.Response{
			Success:  false,
			Messages: err.Error(),
			Results:  nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success:  true,
		Messages: "Reset password success !",
		Results:  nil,
	})
}
