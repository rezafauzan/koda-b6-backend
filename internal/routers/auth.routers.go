package routers

import (
	"rezafauzan/koda-b6-golang/internal/di"
	"rezafauzan/koda-b6-golang/internal/handlers"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	authHandler *handlers.AuthHandler
}

func NewAuthRouters(router *gin.Engine, container *di.Container) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("login", container.AuthHandler.Login)
		authRoutes.POST("register", container.AuthHandler.Register)
	}
}
