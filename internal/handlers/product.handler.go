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

// CreateNewProduct godoc
// @Summary      Create new product
// @Description  Create a new product and store it in database
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateProductRequestDTO  true  "Create Product Request"
// @Success      201      {object}  dto.Response{data=dto.ProductResponseDTO}
// @Failure      400      {object}  dto.Response
// @Failure      500      {object}  dto.Response
// @Router       /products [post]
func (p ProductHandler) CreateNewProduct(ctx *gin.Context) {
	var newProduct dto.CreateProductRequestDTO

	err := ctx.ShouldBindJSON(&newProduct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid request body: " + err.Error(),
			Data:    nil,
		})
		return
	}

	result, err := p.productService.CreateNewProduct(newProduct)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Failed to create product: " + err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "Product created successfully",
		Data:    result,
	})
}

// GetAllProducts godoc
// @Summary      List products
// @Description  Returns all products, optionally filtered by productName query parameter.
// @Tags         products
// @Produce      json
// @Param        productName  query     string  false  "Filter by product name (partial match)"
// @Success      200          {object}  dto.Response{data=[]dto.ProductResponseDTO}
// @Failure      500          {object}  dto.Response
// @Router       /products [get]
func (p ProductHandler) GetAllProducts(ctx *gin.Context) {
	productName := ctx.Query("productName")
	if productName != "" {
		products, err := p.productService.GetAllProductsByName(productName)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, dto.Response{
				Success: false,
				Message: "Failed to create response get all products! : " + err.Error(),
				Data:    nil,
			})
			return
		}

		ctx.JSON(http.StatusOK, dto.Response{
			Success: true,
			Message: "GET all products",
			Data:    products,
		})
		return
	}
	products, err := p.productService.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Failed to create response get all products! : " + err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "GET all products",
		Data:    products,
	})
}

// GetProductById godoc
// @Summary      Get product by ID
// @Description  Returns a single product by its numeric id.
// @Tags         products
// @Produce      json
// @Param        productId  path      int  true  "Product ID"
// @Success      200        {object}  dto.Response{data=dto.ProductResponseDTO}
// @Failure      400        {object}  dto.Response
// @Failure      404        {object}  dto.Response
// @Failure      500        {object}  dto.Response
// @Router       /products/{productId} [get]
func (p ProductHandler) GetProductById(ctx *gin.Context) {
	idParam := ctx.Param("productId")

	productId, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid product id",
			Data:    nil,
		})
		return
	}

	product, err := p.productService.GetProductById(productId)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, dto.Response{
				Success: false,
				Message: "Product not found",
				Data:    nil,
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Failed to get product! : " + err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "GET product by id",
		Data:    product,
	})
}

func (p ProductHandler) UpdateProduct(ctx *gin.Context) {
	var req dto.UpdateProductRequestDTO

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid request body: " + err.Error(),
			Data:    nil,
		})
		return
	}

	id, err := strconv.Atoi(ctx.Param("productId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid product id",
			Data:    nil,
		})
		return
	}

	req.Id = id

	result, err := p.productService.UpdateProduct(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Failed to update product: " + err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Product updated successfully",
		Data:    result,
	})
}
