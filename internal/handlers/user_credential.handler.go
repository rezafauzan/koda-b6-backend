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

func (u UserCredentialHandler) GetAllUserCredentials(ctx *gin.Context) {
	list, err := u.userCredentialService.GetAllUserCredential()
	if err != nil {
		ctx.JSON(http.StatusOK, dto.Response{
			Success:  false,
			Messages: "Failed to create response get all user credentials! : " + err.Error(),
			Results:  nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success:  true,
		Messages: "GET all user credentials",
		Results:  list,
	})
}

func (u UserCredentialHandler) UpdateUserCredential(ctx *gin.Context) {
	var body dto.UpdateUserCredentialDTO
	ctx.ShouldBind(&body)
	updated, err := u.userCredentialService.UpdateUserCredential(body)
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
		Messages: "Update User Credential Success !",
		Results:  updated,
	})
}