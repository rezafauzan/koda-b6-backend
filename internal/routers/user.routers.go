package routers

import (
	"context"
	"net/http"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/lib"
	"rezafauzan/koda-b6-golang/internal/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func NewUserRouters(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("", func(ctx *gin.Context) {
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
		})

		userRoutes.POST("", func(ctx *gin.Context) {
			var newUser dto.UserRegister
			ctx.ShouldBind(&newUser)
			if len(newUser.First_name) < 4 {
				ctx.JSON(http.StatusOK, dto.Response{
					Success:  false,
					Messages: "Failed to create user! : First name length minimum is 4 characters !",
					Results:  nil,
				})
				return
			}
			if len(newUser.Last_name) < 4 {
				ctx.JSON(http.StatusOK, dto.Response{
					Success:  false,
					Messages: "Failed to create user! : Last name length minimum is 4 characters !",
					Results:  nil,
				})
				return
			}
			if !strings.Contains(newUser.Email, "@"){
				ctx.JSON(http.StatusOK, dto.Response{
					Success:  false,
					Messages: "Failed to create user! : Invalid email format !",
					Results:  nil,
				})
				return
			}
			if len(newUser.Phone) < 10 {
				ctx.JSON(http.StatusOK, dto.Response{
					Success:  false,
					Messages: "Failed to create user! : Phone numbers length minimum 10 digits !",
					Results:  nil,
				})
				return
			}
			if len(newUser.Address) < 10 {
				ctx.JSON(http.StatusOK, dto.Response{
					Success:  false,
					Messages: "Failed to create user! : Address length minimum is 10 characters !",
					Results:  nil,
				})
				return
			}
			if len(newUser.Password) < 8 {
				ctx.JSON(http.StatusOK, dto.Response{
					Success:  false,
					Messages: "Failed to create user! : Password too weak minimum length is 8 characters !",
					Results:  nil,
				})
				return
			}
			if newUser.Password_confirm != newUser.Password {
				ctx.JSON(http.StatusOK, dto.Response{
					Success:  false,
					Messages: "Failed to create user! : Password confirmation missmatch !",
					Results:  nil,
				})
				return
			}
			conn, err := lib.DatabaseConnect()
			if err != nil {
				ctx.JSON(http.StatusOK, dto.Response{
					Success:  false,
					Messages: "Failed to connect to database! : " + err.Error(),
					Results:  nil,
				})
				return
			}
			sql := "INSERT INTO users (role_id,verified,created_at,updated_at) VALUES (2, false, $1, $2) RETURNING id, role_id, verified, created_at,updated_at"
			rows, err := conn.Query(context.Background(), sql, time.Now(), time.Now())
			if err != nil {
				ctx.JSON(http.StatusOK, dto.Response{
					Success:  false,
					Messages: "Failed to create new user! : " + err.Error(),
					Results:  nil,
				})
				return
			}

			registeredUser, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
			if err != nil {
				ctx.JSON(http.StatusOK, dto.Response{
					Success:  false,
					Messages: "Failed to create new user! : " + err.Error(),
					Results:  nil,
				})
				return
			}

			sql = "INSERT INTO users_profiles (users_id, user_avatar, first_name, last_name, address, created_at, updated_at) VALUES ($1, https://i.pravatar.cc/400?img=4, $2, $3, $4, $5, $6)"
			commandTag, err := conn.Exec(context.Background(), sql, registeredUser.Id, newUser.First_name, newUser.Last_name, newUser.Address, time.Now(), time.Now())

			if err != nil {
				ctx.JSON(http.StatusOK, dto.Response{
					Success:  false,
					Messages: "Failed to create new user! : " + err.Error(),
					Results:  nil,
				})
				return
			}
			
			sql = "INSERT INTO users_credentials (users_id, email, phone, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"
			commandTag, err = conn.Exec(context.Background(), sql, registeredUser.Id, newUser.Email, newUser.Phone, newUser.Password, time.Now(), time.Now())

			if err != nil {
				ctx.JSON(http.StatusOK, dto.Response{
					Success:  false,
					Messages: "Failed to create new user! : " + err.Error(),
					Results:  nil,
				})
				return
			}
			
			if commandTag.RowsAffected() > 0 {
				ctx.JSON(http.StatusOK, dto.Response{
					Success:  true,
					Messages: "New users successfully",
					Results:  nil,
				})
			}
		})

		userRoutes.PATCH("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, dto.Response{
				Success:  true,
				Messages: "PATCH users",
				Results:  nil,
			})
		})

		userRoutes.DELETE("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, dto.Response{
				Success:  true,
				Messages: "DELETE users",
				Results:  nil,
			})
		})
	}
}
