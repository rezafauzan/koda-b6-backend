package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"os"
	"path/filepath"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/services"
	"strconv"
	"strings"
	"time"

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

func (u UserProfileHandler) UpdateUserAvatar(ctx *gin.Context) {
	userIdRaw, exist := ctx.Get("user_id")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, dto.Response{Success: false, Message: "Invalid or expired token", Data: nil})
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "User avatar file is required", Data: nil})
		return
	}

	contentType := file.Header.Get("Content-Type")
	allowed := map[string]bool{
		"image/png":  true,
		"image/jpg":  true,
		"image/jpeg": true,
		"image/webp": true,
	}
	if !allowed[contentType] {
		ctx.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "Only png, jpg, jpeg, webp files are allowed", Data: nil})
		return
	}

	if file.Size > 10*1024*1024 {
		ctx.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "File too large", Data: nil})
		return
	}

	dstDir := filepath.Join("assets", "img", "user", "avatar")
	if err := os.MkdirAll(dstDir, 0o755); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{Success: false, Message: "Failed to create upload path", Data: nil})
		return
	}

	ext := filepath.Ext(file.Filename)
	if ext == "" {
		ext = ".png"
	}
	token := make([]byte, 8)
	_, _ = rand.Read(token)
	filename := strings.ToLower(hex.EncodeToString(token)) + "_" + strconv.FormatInt(time.Now().UnixNano(), 10) + ext
	dstPath := filepath.Join(dstDir, filename)
	if err := ctx.SaveUploadedFile(file, dstPath); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{Success: false, Message: "Failed to save avatar", Data: nil})
		return
	}

	baseURL := os.Getenv("BASE_URL")
	avatarURL := baseURL + "/assets/img/user/avatar/" + filename
	updated, err := u.userProfileService.UpdateUserProfile(dto.UpdateUserProfileDTO{User_avatar: avatarURL}, userIdRaw.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{Success: false, Message: "Failed update user avatar! " + err.Error(), Data: nil})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{Success: true, Message: "Success update user avatar!", Data: updated})
}
