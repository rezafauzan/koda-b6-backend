package services

import (
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/repository"
)

type ProductService struct {
	productRepository *repository.ProductRepository
}

func NewProductService(productService *repository.ProductRepository) *ProductService{
	return &ProductService{
		productRepository: productService,
	}
}

func (p ProductService) GetAllProducts() ([]dto.ProductResponseDTO, error) {
	products, err := p.productRepository.GetAllProducts()
	
	if err != nil {
		return []dto.ProductResponseDTO{}, err
	}

	var response []dto.ProductResponseDTO

	for _, product := range products {
		modelToDTO := dto.ProductResponseDTO{
			Id: product.Id,
			CategoryId: product.CategoryId,
			Name: product.Name,
			Description: product.Description,
			Price: product.Price,
			Stock: product.Stock,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
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
			Id: product.Id,
			CategoryId: product.CategoryId,
			Name: product.Name,
			Description: product.Description,
			Price: product.Price,
			Stock: product.Stock,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
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