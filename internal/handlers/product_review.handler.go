package handlers

import (
	"net/http"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/services"

	"github.com/gin-gonic/gin"
)

type ProductReviewHandler struct {
	service *services.ProductReviewService
}

func NewProductReviewHandler(service *services.ProductReviewService) *ProductReviewHandler {
	return &ProductReviewHandler{service: service}
}

func (h *ProductReviewHandler) GetLatestReviews(ctx *gin.Context) {
	result, err := h.service.GetLatestReviews()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{Success: false, Message: "Failed to get latest reviews: " + err.Error(), Data: nil})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{Success: true, Message: "Get latest products reviews data", Data: result})
}

func (h *ProductReviewHandler) GetPopularProducts(ctx *gin.Context) {
	result, err := h.service.GetPopularProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{Success: false, Message: "Failed to get popular products: " + err.Error(), Data: nil})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{Success: true, Message: "Get popular products", Data: result})
}
