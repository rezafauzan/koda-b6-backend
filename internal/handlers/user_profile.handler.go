package handlers

import (
	"net/http"
	"rezafauzan/koda-b6-golang/internal/dto"
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

// GetUserProfile godoc
// @Summary      Get my profile
// @Description  Returns the authenticated user's profile (requires Bearer JWT).
// @Tags         profile
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  dto.Response{data=dto.UserProfileResponseDTO}
// @Failure      401  {object}  dto.Response
// @Failure      404  {object}  dto.Response
// @Failure      500  {object}  dto.Response
// @Router       /profile [get]
func (u UserProfileHandler) GetUserProfile(ctx *gin.Context) {
	userId, exist := ctx.Get("user_id")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "Invalid or expired token",
			Data:    nil,
		})
		return
	}

	user, err := u.userProfileService.GetUserProfileByUserId(userId.(int))

	if err != nil {
		msg := err.Error()
		status := http.StatusInternalServerError
		if strings.Contains(msg, "not found") {
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
		Message: "Get User Profile Data!",
		Data:    user,
	})
}

// UpdateUserProfile godoc
// @Summary      Update my profile
// @Description  Updates the authenticated user's profile fields (requires Bearer JWT).
// @Tags         profile
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      dto.UpdateUserProfileDTO  true  "Profile fields"
// @Success      200   {object}  dto.Response{data=dto.UserProfileResponseDTO}
// @Failure      400   {object}  dto.Response
// @Failure      401   {object}  dto.Response
// @Failure      404   {object}  dto.Response
// @Failure      500   {object}  dto.Response
// @Router       /profile [patch]
func (u UserProfileHandler) UpdateUserProfile(ctx *gin.Context) {
	var newData dto.UpdateUserProfileDTO
	err := ctx.ShouldBindJSON(&newData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	userId, exist := ctx.Get("user_id")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "Invalid or expired token",
			Data:    nil,
		})
		return
	}

	updated, err := u.userProfileService.UpdateUserProfile(newData, userId.(int))
	if err != nil {
		msg := err.Error()
		status := http.StatusInternalServerError
		switch {
		case strings.Contains(msg, "minimum"):
			status = http.StatusBadRequest
		case strings.Contains(msg, "not found"):
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
		Message: "Update User Profile Success !",
		Data:    updated,
	})
}
