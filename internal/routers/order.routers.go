package routers

import (
	"rezafauzan/koda-b6-golang/internal/di"
	"rezafauzan/koda-b6-golang/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewOrderRouters(router *gin.Engine, container *di.Container) {
	paymentRoutes := router.Group("/payment")
	{
		paymentRoutes.POST("", middleware.AuthMiddleware(), container.OrderHandler.CreatePayment)
	}
}
