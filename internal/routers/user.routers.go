package routers

import (
	"net/http"
	"rezafauzan/koda-b6-golang/internal/dto"

	"github.com/gin-gonic/gin"
)

func NewUserRouters(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, dto.Response{
				Success: true,
				Messages: "GET users",
				Results: nil,
			})
		})
		userRoutes.POST("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, dto.Response{
				Success: true,
				Messages: "POST users",
				Results: nil,
			})
		})
		userRoutes.PATCH("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, dto.Response{
				Success: true,
				Messages: "PATCH users",
				Results: nil,
			})
		})
		userRoutes.DELETE("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, dto.Response{
				Success: true,
				Messages: "DELETE users",
				Results: nil,
			})
		})
	}
}
