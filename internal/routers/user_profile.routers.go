package routers

import (
	"rezafauzan/koda-b6-golang/internal/di"
	"rezafauzan/koda-b6-golang/internal/handlers"
	"rezafauzan/koda-b6-golang/internal/middleware"

	"github.com/gin-gonic/gin"
)

type UserProfileRouter struct {
	userProfileHandler *handlers.UserProfileHandler
}

func NewUserProfileRouters(router *gin.Engine, container *di.Container) {
	userProfileRoutes := router.Group("/profile")
	{
		userProfileRoutes.GET("", middleware.AuthMiddleware(), container.UserProfileHandler.GetUserProfile)

		userProfileRoutes.PATCH("", middleware.AuthMiddleware(), container.UserProfileHandler.UpdateUserProfile)
	}
}
