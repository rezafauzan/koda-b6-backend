package handlers

import (
	"context"
	"net/http"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/lib"
	"rezafauzan/koda-b6-golang/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandlers() (*UserHandler){
	return &UserHandler{
		userService: &services.UserService{},
	}
}

func (u UserHandler) GetAllUsers(ctx *gin.Context) {
	conn, err := lib.DatabaseConnect()
	if err != nil {
		ctx.JSON(http.StatusOK, dto.Response{
			Success:  false,
			Messages: "Failed to connect to database! : " + err.Error(),
			Results:  nil,
		})
		return
	}

	sql := `SELECT users.id, user_profiles.user_avatar, user_profiles.first_name, user_profiles.last_name, user_credentials.email, user_credentials.phone, user_profiles.address, users.verified, roles.role_name, users.created_at, users.updated_at FROM users JOIN roles ON roles.id = users.role_id JOIN user_profiles ON user_profiles.user_id = users.id JOIN user_credentials ON user_credentials.user_id = users.id`
	rows, err := conn.Query(context.Background(), sql)
	if err != nil {
		ctx.JSON(http.StatusOK, dto.Response{
			Success:  false,
			Messages: "Failed to get all users! : " + err.Error(),
			Results:  nil,
		})
		return
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.User])
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
		Messages: "GET users",
		Results:  users,
	})
}

func (u UserHandler) AddNewUser(ctx *gin.Context) {
	var newUser dto.UserRegister
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
		Success:  false,
		Messages: "Register Success!",
		Results:  newUser,
	})
}
