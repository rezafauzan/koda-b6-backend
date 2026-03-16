package routers

import (
	"net/http"
	"rezafauzan/koda-b6-golang/internal/di"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/handlers"

	"github.com/gin-gonic/gin"
)
type UserRouter struct{
	userHandler *handlers.UserHandler
}

func NewUserRouters(router *gin.Engine, container *di.Container) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("", container.UserHandler.GetAllUsers)

		userRoutes.POST("", container.UserHandler.AddNewUser)

		userRoutes.PATCH("", container.UserHandler.UpdateUserProfiles)

		userRoutes.DELETE("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, dto.Response{
				Success:  true,
				Messages: "DELETE users",
				Results:  nil,
			})
		})
	}
}
