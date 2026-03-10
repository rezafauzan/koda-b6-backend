package routers

import (
	"context"
	"net/http"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/lib"
	"rezafauzan/koda-b6-golang/internal/models"

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
			} else {
				sql := "SELECT * FROM users"
				rows, err := conn.Query(context.Background(), sql)
				if err != nil {
					ctx.JSON(http.StatusOK, dto.Response{
						Success:  false,
						Messages: "Failed to get all users! : " + err.Error(),
						Results:  nil,
					})
				} else {
					users, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.User])
					if err != nil {
						ctx.JSON(http.StatusOK, dto.Response{
							Success:  false,
							Messages: "Failed to create response get all users! : " + err.Error(),
							Results:  nil,
						})
					} else {
						ctx.JSON(http.StatusOK, dto.Response{
							Success:  true,
							Messages: "GET users",
							Results:  users,
						})
					}
				}
			}
		})

		userRoutes.POST("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, dto.Response{
				Success:  true,
				Messages: "POST users",
				Results:  nil,
			})
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
