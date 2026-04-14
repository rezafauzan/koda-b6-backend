package handlers

import (
	"errors"
	"net/http"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/services"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetAllUsers godoc
// @Summary      List users
// @Description  Returns all registered users with profile and role information.
// @Tags         users
// @Produce      json
// @Success      200  {object}  dto.Response{data=[]dto.UserResponseDTO}
// @Failure      500  {object}  dto.Response
// @Router       /users [get]
func (u UserHandler) GetAllUsers(ctx *gin.Context) {
	users, err := u.userService.GetAllUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Failed to create response get all users! : " + err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "GET all users",
		Data:    users,
	})
}

// CreateNewUser godoc
// @Summary      Register user
// @Description  Creates a new user account with profile and credentials.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        body  body      dto.CreateUserDTO  true  "Registration payload"
// @Success      201   {object}  dto.Response{data=dto.CreateUserDTO}
// @Failure      400   {object}  dto.Response
// @Failure      409   {object}  dto.Response
// @Failure      500   {object}  dto.Response
// @Router       /users [post]
func (u UserHandler) CreateNewUser(ctx *gin.Context) {
	var newUser dto.CreateUserDTO
	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	newUser, err = u.userService.CreateNewUser(newUser)
	if err != nil {
		msg := err.Error()
		status := http.StatusInternalServerError
		switch {
		case strings.Contains(msg, "minimum") || strings.Contains(msg, "Invalid email") || strings.Contains(msg, "missmatch"):
			status = http.StatusBadRequest
		case strings.Contains(msg, "allready used"):
			status = http.StatusConflict
		}
		ctx.JSON(status, dto.Response{
			Success: false,
			Message: msg,
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "Register Success!",
		Data:    newUser,
	})
}

// UpdateUserProfiles godoc
// @Summary      Update user profile
// @Description  Updates profile fields for a user identified in the body payload.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        body  body      dto.UpdateUserProfileDTO  true  "Profile update payload"
// @Success      200   {object}  dto.Response{data=dto.UserResponseDTO}
// @Failure      400   {object}  dto.Response
// @Failure      404   {object}  dto.Response
// @Failure      500   {object}  dto.Response
// @Router       /users [patch]
func (u UserHandler) UpdateUserProfiles(ctx *gin.Context) {
	var newUser dto.UpdateUserProfileDTO
	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid user id !",
			Data:    nil,
		})
		return
	}
	newUser.Id = id

	updatedUser, err := u.userService.UpdateUserProfiles(newUser)
	if err != nil {
		msg := err.Error()
		status := http.StatusInternalServerError
		if errors.Is(err, pgx.ErrNoRows) || strings.Contains(msg, "no rows") {
			status = http.StatusNotFound
		} else if strings.Contains(msg, "minimum") {
			status = http.StatusBadRequest
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
		Message: "Update Users Success !",
		Data:    updatedUser,
	})
}

// DeleteUser godoc
// @Summary      Delete user
// @Description  Permanently removes a user and related records by id.
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      204  "No Content"
// @Failure      400  {object}  dto.Response
// @Failure      404  {object}  dto.Response
// @Failure      500  {object}  dto.Response
// @Router       /users/{id} [delete]
func (u UserHandler) DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid user id !",
			Data:    nil,
		})
		return
	}
	deleted, err := u.userService.DeleteUser(id)
	if err != nil {
		msg := err.Error()
		status := http.StatusInternalServerError
		if strings.Contains(msg, "User not found") {
			status = http.StatusNotFound
		}
		ctx.JSON(status, dto.Response{
			Success: false,
			Message: "Delete user failed : " + msg,
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Delete user success!",
		Data:    deleted,
	})
}
