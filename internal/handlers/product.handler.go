package handlers

import (
	"net/http"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type ProductHandler struct {
	productService *services.ProductService
}

func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (p ProductHandler) GetAllProducts(ctx *gin.Context) {
	productName := ctx.Query("productName")
	if productName != "" {
		products, err := p.productService.GetAllProductsByName(productName)
		if err != nil {
			ctx.JSON(http.StatusOK, dto.Response{
				Success:  false,
				Messages: "Failed to create response get all products! : " + err.Error(),
				Results:  nil,
			})
			return
		}

		ctx.JSON(http.StatusOK, dto.Response{
			Success:  true,
			Messages: "GET all products",
			Results:  products,
		})
		return
	}
	products, err := p.productService.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusOK, dto.Response{
			Success:  false,
			Messages: "Failed to create response get all products! : " + err.Error(),
			Results:  nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success:  true,
		Messages: "GET all products",
		Results:  products,
	})
}

func (p ProductHandler) GetProductById(ctx *gin.Context) {
	idParam := ctx.Param("productId")

	productId, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success:  false,
			Messages: "Invalid product id",
			Results:  nil,
		})
		return
	}

	product, err := p.productService.GetProductById(productId)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, dto.Response{
				Success:  false,
				Messages: "Product not found",
				Results:  nil,
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success:  false,
			Messages: "Failed to get product! : " + err.Error(),
			Results:  nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success:  true,
		Messages: "GET product by id",
		Results:  product,
	})
}
