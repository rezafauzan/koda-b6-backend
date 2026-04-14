package routers

import (
	"rezafauzan/koda-b6-golang/internal/di"

	"github.com/gin-gonic/gin"
)

func NewProductReviewRouters(router *gin.Engine, container *di.Container) {
	router.GET("/recommended-products", container.ProductReviewHandler.GetPopularProducts)
	router.GET("/reviews", container.ProductReviewHandler.GetLatestReviews)
}
