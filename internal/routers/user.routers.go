package routers

import (
	"rezafauzan/koda-b6-golang/internal/di"
	"rezafauzan/koda-b6-golang/internal/handlers"
	"rezafauzan/koda-b6-golang/internal/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	userHandler *handlers.UserHandler
}

func NewUserRouters(router *gin.Engine, container *di.Container) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("", container.UserHandler.GetAllUsers)

		userRoutes.GET("logged-in", middleware.AuthMiddleware() ,container.UserHandler.GetLoggedInUser)

		userRoutes.POST("", container.UserHandler.CreateNewUser)

		userRoutes.PATCH("", container.UserHandler.UpdateUserProfiles)

		userRoutes.DELETE(":id", container.UserHandler.DeleteUser)
	}
}
