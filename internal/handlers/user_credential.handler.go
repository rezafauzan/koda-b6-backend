package handlers

import (
	"net/http"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/services"

	"github.com/gin-gonic/gin"
)

type UserCredentialHandler struct {
	userCredentialService *services.UserCredentialService
}

func NewUserCredentialHandler(userCredentialService *services.UserCredentialService) *UserCredentialHandler {
	return &UserCredentialHandler{
		userCredentialService: userCredentialService,
	}
}

func (u UserCredentialHandler) GetUserCredentialsByUserId(ctx *gin.Context) {
	userId, exist := ctx.Get("user_id")
	if !exist {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success:  false,
			Messages: "Invalid or expired token",
			Results:  nil,
		})
		return
	}
	userCredentials, err := u.userCredentialService.GetUserCredentialByUserId(userId.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success:  false,
			Messages: "Failed to create response get user credentials! : " + err.Error(),
			Results:  nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success:  true,
		Messages: "GET user credentials",
		Results:  userCredentials,
	})
}

func (u UserCredentialHandler) UpdateUserCredential(ctx *gin.Context) {
	var body dto.UpdateUserCredentialDTO
	ctx.ShouldBind(&body)
	updated, err := u.userCredentialService.UpdateUserCredential(body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success:  false,
			Messages: err.Error(),
			Results:  nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success:  true,
		Messages: "Update User Credential Success !",
		Results:  updated,
	})
}
