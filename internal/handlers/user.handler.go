package handlers

import (
	"net/http"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (u UserHandler) GetAllUsers(ctx *gin.Context) {
	users, err := u.userService.GetAllUser()
	if err != nil {
		ctx.JSON(http.StatusOK, dto.Response{
			Success:  false,
			Messages: "Failed to create response get all users! : " + err.Error(),
			Results:  nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success:  true,
		Messages: "GET all users",
		Results:  users,
	})
}

func (u UserHandler) AddNewUser(ctx *gin.Context) {
	var newUser *dto.CreateUserDTO
	ctx.ShouldBind(&newUser)
	newUser, err := u.userService.AddNewUser(newUser)
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
		Messages: "Register Success!",
		Results:  newUser,
	})
}

func (u UserHandler) UpdateUserProfiles(ctx *gin.Context) {
	var newUser dto.UpdateUserProfileDTO
	ctx.ShouldBind(&newUser)
	updatedUser, err := u.userService.UpdateUserProfiles(newUser)
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
		Messages: "Update Users Success !",
		Results:  updatedUser,
	})
}

func (u UserHandler) DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success:  false,
			Messages: "Invalid user id !",
			Results:  nil,
		})
		return
	}
	deletedUser, err := u.userService.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusOK, dto.Response{
			Success:  false,
			Messages: "Delete user failed : " + err.Error(),
			Results:  nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success:  true,
		Messages: "Delete Users Success !",
		Results:  deletedUser,
	})
}
