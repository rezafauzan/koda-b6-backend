package handlers

import (
	"net/http"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/services"
	"strings"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService *services.OrderService
}

func NewOrderHandler(orderService *services.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) CreatePayment(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, dto.Response{Success: false, Message: "Invalid user session", Data: nil})
		return
	}

	var req dto.CreatePaymentRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: err.Error(), Data: nil})
		return
	}

	result, err := h.orderService.CreatePayment(userID.(int), req)
	if err != nil {
		status := http.StatusInternalServerError
		msg := err.Error()
		if strings.Contains(msg, "Checkout failed") || strings.Contains(msg, "Cart is empty") {
			status = http.StatusBadRequest
		}
		if strings.Contains(msg, "Invalid user session") {
			status = http.StatusUnauthorized
		}
		ctx.JSON(status, dto.Response{Success: false, Message: msg, Data: nil})
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{Success: true, Message: "Checkout success!", Data: result})
}
