package routers

import (
	"rezafauzan/koda-b6-golang/internal/di"
	"rezafauzan/koda-b6-golang/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewProductRouter(router *gin.Engine, container *di.Container){
	productRoutes := router.Group("/products")
	{
		productRoutes.POST("", middleware.AuthMiddleware(), middleware.RBAC("admin"), container.ProductHandler.CreateNewProduct)
		productRoutes.GET("", container.ProductHandler.GetAllProducts)
		productRoutes.GET("/category/:category_id", container.ProductHandler.GetProductByCategoryId)
		productRoutes.GET("/:productId", container.ProductHandler.GetProductById)
		productRoutes.PUT("/:productId", middleware.AuthMiddleware(), middleware.RBAC("admin"), container.ProductHandler.UpdateProduct)
		productRoutes.DELETE("/:productId", middleware.AuthMiddleware(), middleware.RBAC("admin"), container.ProductHandler.DeleteProduct)
	}
}