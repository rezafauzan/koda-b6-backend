package routers

import (
	"rezafauzan/koda-b6-golang/internal/di"
	"rezafauzan/koda-b6-golang/internal/handlers"

	"github.com/gin-gonic/gin"
)

type RoleRouter struct {
	roleHandler *handlers.RoleHandler
}

func NewRoleRouters(router *gin.Engine, container *di.Container) {
	roleRoutes := router.Group("/roles")
	{
		roleRoutes.GET("", container.RoleHandler.GetAllRoles)
		roleRoutes.POST("", container.RoleHandler.CreateNewRole)
		roleRoutes.PATCH("", container.RoleHandler.UpdateRole)
		roleRoutes.DELETE(":id", container.RoleHandler.DeleteRole)
	}
}