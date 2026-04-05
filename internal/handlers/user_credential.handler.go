package handlers

import (
	"net/http"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/services"
	"strings"

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

// GetUserCredentialsByUserId godoc
// @Summary      Get my credentials
// @Description  Returns email and phone for the authenticated user (requires Bearer JWT).
// @Tags         user-credentials
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  dto.Response{data=dto.UserCredentialResponseWithoutPasswordDTO}
// @Failure      401  {object}  dto.Response
// @Failure      404  {object}  dto.Response
// @Failure      500  {object}  dto.Response
// @Router       /user-credentials [get]
func (u UserCredentialHandler) GetUserCredentialsByUserId(ctx *gin.Context) {
	userId, exist := ctx.Get("user_id")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "Invalid or expired token",
			Data:    nil,
		})
		return
	}
	userCredentials, err := u.userCredentialService.GetUserCredentialByUserId(userId.(int))
	if err != nil {
		msg := err.Error()
		status := http.StatusInternalServerError
		if strings.Contains(msg, "not found") {
			status = http.StatusNotFound
		}
		ctx.JSON(status, dto.Response{
			Success: false,
			Message: "Failed to create response get user credentials! : " + msg,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "GET user credentials",
		Data:    userCredentials,
	})
}

// UpdateUserCredential godoc
// @Summary      Update credentials
// @Description  Updates email, phone, or password for the account identified in the body (requires Bearer JWT).
// @Tags         user-credentials
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      dto.UpdateUserCredentialDTO  true  "Credential update"
// @Success      200   {object}  dto.Response{data=dto.UserCredentialResponseDTO}
// @Failure      400   {object}  dto.Response
// @Failure      404   {object}  dto.Response
// @Failure      500   {object}  dto.Response
// @Router       /user-credentials [patch]
func (u UserCredentialHandler) UpdateUserCredential(ctx *gin.Context) {
	var body dto.UpdateUserCredentialDTO
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	updated, err := u.userCredentialService.UpdateUserCredential(body)
	if err != nil {
		msg := err.Error()
		status := http.StatusInternalServerError
		switch {
		case strings.Contains(msg, "minimum") || strings.Contains(msg, "Invalid email") || strings.Contains(msg, "not matched"):
			status = http.StatusBadRequest
		case strings.Contains(strings.ToLower(msg), "no rows") || strings.Contains(msg, "not found"):
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
		Message: "Update User Credential Success !",
		Data:    updated,
	})
}
