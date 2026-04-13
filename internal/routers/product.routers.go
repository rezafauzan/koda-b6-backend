package routers

import (
	"rezafauzan/koda-b6-golang/internal/di"

	"github.com/gin-gonic/gin"
)

func NewProductRouter(router *gin.Engine, container *di.Container){
	productRoutes := router.Group("/products")
	{
		productRoutes.POST("", container.ProductHandler.CreateNewProduct)
		productRoutes.GET("", container.ProductHandler.GetAllProducts)
		productRoutes.GET("/:productId", container.ProductHandler.GetProductById)
	}
}