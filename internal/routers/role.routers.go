package routers

import (
	"rezafauzan/koda-b6-golang/internal/di"
	"rezafauzan/koda-b6-golang/internal/handlers"
	"rezafauzan/koda-b6-golang/internal/middleware"

	"github.com/gin-gonic/gin"
)

type RoleRouter struct {
	roleHandler *handlers.RoleHandler
}

func NewRoleRouters(router *gin.Engine, container *di.Container) {
	roleRoutes := router.Group("/role")
	{
		roleRoutes.GET("", middleware.AuthMiddleware(), middleware.RBAC("admin"), container.RoleHandler.GetAllRoles)
		roleRoutes.POST("", middleware.AuthMiddleware(), middleware.RBAC("admin"), container.RoleHandler.CreateNewRole)
		roleRoutes.GET("/:name", middleware.AuthMiddleware(), middleware.RBAC("admin"), container.RoleHandler.GetRoleByName)
		roleRoutes.PUT("/:id", middleware.AuthMiddleware(), middleware.RBAC("admin"), container.RoleHandler.UpdateRole)
		roleRoutes.DELETE("/:id", middleware.AuthMiddleware(), middleware.RBAC("admin"), container.RoleHandler.DeleteRole)
	}
}