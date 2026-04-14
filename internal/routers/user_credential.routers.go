package routers

import (
	"rezafauzan/koda-b6-golang/internal/di"
	"rezafauzan/koda-b6-golang/internal/handlers"
	"rezafauzan/koda-b6-golang/internal/middleware"

	"github.com/gin-gonic/gin"
)

type UserCredentialRouter struct {
	userCredentialHandler *handlers.UserCredentialHandler
}

func NewUserCredentialRouters(router *gin.Engine, container *di.Container) {
	userCredentialRoutes := router.Group("/credentials")
	{
		userCredentialRoutes.GET("", middleware.AuthMiddleware(), container.UserCredentialHandler.GetUserCredentialsByUserId)
		userCredentialRoutes.PATCH("", middleware.AuthMiddleware(), container.UserCredentialHandler.UpdateUserCredential)
	}
}
