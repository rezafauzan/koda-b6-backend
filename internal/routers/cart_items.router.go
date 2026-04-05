package routers

import (
	"rezafauzan/koda-b6-golang/internal/di"
	"rezafauzan/koda-b6-golang/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewCartItemRouters(router *gin.Engine, container *di.Container) {
	cartItemRoutes := router.Group("/cart")
	{
		cartItemRoutes.GET("", middleware.AuthMiddleware(), container.CartItemHandler.GetCartItemsByCartId)
		cartItemRoutes.POST("", middleware.AuthMiddleware(), container.CartItemHandler.AddItem)
		cartItemRoutes.DELETE("/:id", middleware.AuthMiddleware(), container.CartItemHandler.DeleteItem)
	}
}