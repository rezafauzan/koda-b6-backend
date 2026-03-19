package handlers

import (
	"net/http"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/services"
	"strconv"

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

func (u UserProfileHandler) GetAllUserProfiles(ctx *gin.Context) {
	list, err := u.userProfileService.GetAllUserProfile()
	if err != nil {
		ctx.JSON(http.StatusOK, dto.Response{
			Success:  false,
			Messages: "Failed to create response get all user profiles! : " + err.Error(),
			Results:  nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success:  true,
		Messages: "GET all user profiles",
		Results:  list,
	})
}

func (u UserProfileHandler) CreateNewUserProfile(ctx *gin.Context) {
	var body dto.CreateUserProfileDTO
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success:  false,
			Messages: err.Error(),
			Results:  nil,
		})
		return
	}
	result, err := u.userProfileService.CreateNewUserProfile(body)
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
		Messages: "Create User Profile Success !",
		Results:  result,
	})
}

func (u UserProfileHandler) UpdateUserProfileEntity(ctx *gin.Context) {
	var body dto.UpdateUserProfileEntityDTO
	ctx.ShouldBind(&body)
	updated, err := u.userProfileService.UpdateUserProfileEntity(body)
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

func (u UserProfileHandler) DeleteUserProfile(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success:  false,
			Messages: "Invalid user profile id !",
			Results:  nil,
		})
		return
	}
	deleted, err := u.userProfileService.DeleteUserProfile(id)
	if err != nil {
		ctx.JSON(http.StatusOK, dto.Response{
			Success:  false,
			Messages: "Delete user profile failed : " + err.Error(),
			Results:  nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success:  true,
		Messages: "Delete User Profile Success !",
		Results:  deleted,
	})
}
