package handlers

import (
	"net/http"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/services"

	"github.com/gin-gonic/gin"
)

type UserProfileHandler struct {
	userProfileService *services.UserProfileService
}

func NewUserProfileHandler(userProfileService *services.UserProfileService) *UserProfileHandler {
	return &UserProfileHandler{
		userProfileService: userProfileService,
	}
}

func (u UserProfileHandler) GetUserProfile(ctx *gin.Context) {
	var newData dto.UpdateUserProfileDTO
	ctx.ShouldBind(&newData)
	userId, exist := ctx.Get("user_id")
	if !exist {
		ctx.JSON(http.StatusOK, dto.Response{
			Success:  false,
			Messages: "Invalid or expired token",
			Results:  nil,
		})
		return
	}

	user, err := u.userProfileService.GetUserProfileByUserId(userId.(int))

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
		Messages: "Get User Profile Data!",
		Results:  user,
	})
}

func (u UserProfileHandler) UpdateUserProfile(ctx *gin.Context) {
	var newData dto.UpdateUserProfileDTO
	ctx.ShouldBind(&newData)
	userId, exist := ctx.Get("user_id")
	if !exist {
		ctx.JSON(http.StatusOK, dto.Response{
			Success:  false,
			Messages: "Invalid or expired token",
			Results:  nil,
		})
		return
	}

	updated, err := u.userProfileService.UpdateUserProfile(newData, userId.(int))
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
		Messages: "Update User Profile Success !",
		Results:  updated,
	})
}