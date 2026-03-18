package routers

import (
	"rezafauzan/koda-b6-golang/internal/di"
	"rezafauzan/koda-b6-golang/internal/handlers"

	"github.com/gin-gonic/gin"
)

type UserCredentialRouter struct {
	userCredentialHandler *handlers.UserCredentialHandler
}

func NewUserCredentialRouters(router *gin.Engine, container *di.Container) {
	userCredentialRoutes := router.Group("/user-credentials")
	{
		userCredentialRoutes.GET("", container.UserCredentialHandler.GetAllUserCredentials)

		userCredentialRoutes.PATCH("", container.UserCredentialHandler.UpdateUserCredential)
	}
}
