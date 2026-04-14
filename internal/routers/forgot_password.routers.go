package routers

import (
	"rezafauzan/koda-b6-golang/internal/di"
	"rezafauzan/koda-b6-golang/internal/handlers"

	"github.com/gin-gonic/gin"
)

type ForgotPasswordRouter struct {
	forgotPasswordHandler *handlers.ForgotPasswordHandler
}

func NewForgotPasswordRouters(router *gin.Engine, container *di.Container) {
	forgotPasswordRoutes := router.Group("/forgot-password")
	{
		forgotPasswordRoutes.POST("/request", container.ForgotPasswordHandler.RequestForgotPassword)
		forgotPasswordRoutes.POST("/reset", container.ForgotPasswordHandler.ResetPassword)
	}
}
