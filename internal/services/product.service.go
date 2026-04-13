package services

import (
	"errors"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/models"
	"rezafauzan/koda-b6-golang/internal/repository"
	"strings"
)

type ProductService struct {
	productRepository *repository.ProductRepository
}

func NewProductService(productService *repository.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productService,
	}
}

func (p ProductService) CreateNewProduct(newProduct dto.CreateProductRequestDTO) (dto.ProductResponseDTO, error) {

	if strings.TrimSpace(newProduct.Name) == "" {
		return dto.ProductResponseDTO{}, errors.New("product name is required")
	}

	if newProduct.CategoryId <= 0 {
		return dto.ProductResponseDTO{}, errors.New("category_id must be valid")
	}

	if strings.TrimSpace(newProduct.Description) == "" {
		return dto.ProductResponseDTO{}, errors.New("description is required")
	}

	if newProduct.Price <= 0 {
		return dto.ProductResponseDTO{}, errors.New("price must be greater than 0")
	}

	if newProduct.Stock < 0 {
		return dto.ProductResponseDTO{}, errors.New("stock cannot be negative")
	}

	modeledNewProduct := models.Product{
		Name:        newProduct.Name,
		CategoryId:  newProduct.CategoryId,
		Description: newProduct.Description,
		Price:       newProduct.Price,
		Stock:       newProduct.Stock,
	}
	registeredProduct, err := p.productRepository.CreateProduct(&modeledNewProduct)
	if err != nil {
		return dto.ProductResponseDTO{}, err
	}

	response := dto.ProductResponseDTO{
		Id:          registeredProduct.Id,
		Name:        registeredProduct.Name,
		CategoryId:  registeredProduct.CategoryId,
		Description: registeredProduct.Description,
		Price:       registeredProduct.Price,
		Stock:       registeredProduct.Stock,
		CreatedAt:   registeredProduct.CreatedAt,
		UpdatedAt:   registeredProduct.UpdatedAt,
	}

	return response, nil
}

func (p ProductService) GetAllProducts() ([]dto.ProductResponseDTO, error) {
	products, err := p.productRepository.GetAllProducts()

	if err != nil {
		return []dto.ProductResponseDTO{}, err
	}

	var response []dto.ProductResponseDTO

	for _, product := range products {
		modelToDTO := dto.ProductResponseDTO{
			Id:          product.Id,
			CategoryId:  product.CategoryId,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		}

		response = append(response, modelToDTO)
	}

	return response, nil
}

func (p ProductService) GetAllProductsByName(productName string) ([]dto.ProductResponseDTO, error) {
	products, err := p.productRepository.GetAllProductsByName(productName)

	if err != nil {
		return []dto.ProductResponseDTO{}, err
	}

	var response []dto.ProductResponseDTO

	for _, product := range products {
		modelToDTO := dto.ProductResponseDTO{
			Id:          product.Id,
			CategoryId:  product.CategoryId,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		}

		response = append(response, modelToDTO)
	}

	return response, nil
}

func (p ProductService) GetProductById(productId int) (dto.ProductResponseDTO, error) {
	product, err := p.productRepository.GetProductById(productId)
	if err != nil {
		return dto.ProductResponseDTO{}, err
	}

	result := dto.ProductResponseDTO{
		Id:          product.Id,
		CategoryId:  product.CategoryId,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}

	return result, nil
}
