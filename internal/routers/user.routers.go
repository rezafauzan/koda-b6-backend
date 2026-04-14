package routers

import (
	"rezafauzan/koda-b6-golang/internal/di"
	"rezafauzan/koda-b6-golang/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewUserRouters(router *gin.Engine, container *di.Container) {
	userRoutes := router.Group("/admin/users")
	{
		userRoutes.GET("", middleware.AuthMiddleware(), middleware.RBAC("admin"), container.UserHandler.GetAllUsers)
		userRoutes.POST("", middleware.AuthMiddleware(), middleware.RBAC("admin"), container.UserHandler.CreateNewUser)
		userRoutes.PATCH("/:id", middleware.AuthMiddleware(), middleware.RBAC("admin"), container.UserHandler.UpdateUserProfiles)
		userRoutes.DELETE("/:id", middleware.AuthMiddleware(), middleware.RBAC("admin"), container.UserHandler.DeleteUser)
	}
}
