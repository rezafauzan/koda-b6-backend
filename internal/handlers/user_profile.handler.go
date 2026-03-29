package handlers

import (
	"net/http"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/lib"
	"rezafauzan/koda-b6-golang/internal/services"
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
	var newData dto.UpdateUserProfileDTO
	ctx.ShouldBind(&newData)
	user_id, exist := ctx.Get("user_id")
	if !exist {
		ctx.JSON(http.StatusOK, dto.Response{
			Success:  false,
			Messages: "Invalid or expired token",
			Results:  nil,
		})
		return
	}

	updated, err := u.userProfileService.UpdateUserProfile(newData, user_id.(int))
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