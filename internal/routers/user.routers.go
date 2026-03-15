package routers

import (
	"net/http"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/handlers"

	"github.com/gin-gonic/gin"
)

func NewUserRouters(router *gin.Engine) {
	userHandler := handlers.NewUserHandlers()
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("", userHandler.GetAllUsers)

		userRoutes.POST("", userHandler.AddNewUser)

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
