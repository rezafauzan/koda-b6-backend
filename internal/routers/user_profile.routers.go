package routers

import (
	"rezafauzan/koda-b6-golang/internal/di"
	"rezafauzan/koda-b6-golang/internal/handlers"

	"github.com/gin-gonic/gin"
)

type UserProfileRouter struct {
	userProfileHandler *handlers.UserProfileHandler
}

func NewUserProfileRouters(router *gin.Engine, container *di.Container) {
	userProfileRoutes := router.Group("/user-profiles")
	{
		userProfileRoutes.GET("", container.UserProfileHandler.GetAllUserProfiles)

		userProfileRoutes.POST("", container.UserProfileHandler.CreateNewUserProfile)

		userProfileRoutes.PATCH("", container.UserProfileHandler.UpdateUserProfileEntity)

		userProfileRoutes.DELETE(":id", container.UserProfileHandler.DeleteUserProfile)
	}
}
