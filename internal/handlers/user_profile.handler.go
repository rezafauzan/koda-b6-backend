package handlers

import (
	"net/http"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/lib"
	"rezafauzan/koda-b6-golang/internal/services"
	"strconv"
	"strings"

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

func (u UserProfileHandler) GetUserProfile(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" {
		ctx.JSON(http.StatusBadRequest,
			dto.Response{
				Success:  false,
				Messages: "Error missing token",
				Results:  nil,
			})
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	token = strings.TrimSpace(token)

	claims, err := lib.VerifyJWT(token)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success:  false,
			Messages: err.Error(),
			Results:  nil,
		})
		return
	}

	user, err := u.userProfileService.GetUserProfileByUserId(claims.User_id)

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
