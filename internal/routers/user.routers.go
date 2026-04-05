package routers

import (
	"rezafauzan/koda-b6-golang/internal/di"

	"github.com/gin-gonic/gin"
)

func NewUserRouters(router *gin.Engine, container *di.Container) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("", container.UserHandler.GetAllUsers)
		userRoutes.POST("", container.UserHandler.CreateNewUser)
		userRoutes.PATCH("", container.UserHandler.UpdateUserProfiles)
		userRoutes.DELETE(":id", container.UserHandler.DeleteUser)
	}
}
