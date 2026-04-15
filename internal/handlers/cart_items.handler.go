package handlers

import (
	"net/http"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type CartItemHandler struct {
	cartItemService *services.CartItemService
}

func NewCartItemHandler(cartItemService *services.CartItemService) *CartItemHandler {
	return &CartItemHandler{cartItemService: cartItemService}
}

// AddCartItem godoc
// @Summary      Add item to cart
// @Description  Add product item into cart
// @Tags         cart-items
// @Accept       json
// @Produce      json
// @Success      201  {object}  dto.Response{data=models.CartItem}
// @Failure      400  {object}  dto.Response
// @Failure      500  {object}  dto.Response
// @Router       /cart-items [post]
func (c CartItemHandler) AddItem(ctx *gin.Context) {
	cartId, exist := ctx.Get("cart_id")
	if !exist {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid token",
			Data:    nil,
		})
		return
	}

	var newCartItemData dto.CreateCartItemRequestDTO

	err := ctx.ShouldBindJSON(&newCartItemData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}
	newCartItemData.CartId = cartId.(int)

	result, err := c.cartItemService.AddItem(newCartItemData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Failed to create cart item! : " + err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "ADD cart item",
		Data:    result,
	})
}

// GetCartItemsByCartId godoc
// @Summary      Get cart items
// @Description  Get all items by cart id
// @Tags         cart-items
// @Produce      json
// @Param        cartId  path  int  true  "Cart ID"
// @Success      200  {object}  dto.Response{data=[]models.CartItem}
// @Failure      400  {object}  dto.Response
// @Failure      500  {object}  dto.Response
// @Router       /cart-items/cart/{cartId} [get]
func (c CartItemHandler) GetCartItemsByCartId(ctx *gin.Context) {
	cartId, exist := ctx.Get("cart_id")
	if !exist {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid cart id",
			Data:    nil,
		})
		return
	}

	result, err := c.cartItemService.GetCartItemsByCartId(cartId.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Failed to get cart items! : " + err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "GET cart items by cart id",
		Data:    result,
	})
}

// DeleteCartItem godoc
// @Summary      Delete cart item
// @Description  Delete cart item by id
// @Tags         cart-items
// @Produce      json
// @Param        id  path  int  true  "Cart Item ID"
// @Success      200  {object}  dto.Response{data=models.CartItem}
// @Failure      400  {object}  dto.Response
// @Failure      404  {object}  dto.Response
// @Failure      500  {object}  dto.Response
// @Router       /cart-items/{id} [delete]
func (c CartItemHandler) DeleteItem(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid cart item id",
			Data:    nil,
		})
		return
	}

	result, err := c.cartItemService.DeleteItem(id)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, dto.Response{
				Success: false,
				Message: "Cart item not found",
				Data:    nil,
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Failed to delete cart item! : " + err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "DELETE cart item",
		Data:    result,
	})
}
