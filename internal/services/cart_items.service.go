package services

import (
	"errors"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/models"
	"rezafauzan/koda-b6-golang/internal/repository"
)

type CartItemService struct {
	repo *repository.CartItemRepository
}

func NewCartItemService(repo *repository.CartItemRepository) *CartItemService {
	return &CartItemService{repo: repo}
}

func (s *CartItemService) AddItem(newCartItemData dto.CreateCartItemRequestDTO) (models.CartItem, error) {
	if newCartItemData.ProductId <= 0 {
		return models.CartItem{}, errors.New("Failed to add item! : Product ID is required !")
	}

	if newCartItemData.SizeId <= 0 {
		return models.CartItem{}, errors.New("Failed to add item! : Size is required !")
	}

	if newCartItemData.VariantId <= 0 {
		return models.CartItem{}, errors.New("Failed to add item! : Variant is required !")
	}

	if newCartItemData.Quantity <= 0 {
		return models.CartItem{}, errors.New("Failed to add item! : Quantity minimum is 1 !")
	}

	result, err := s.repo.AddItem(newCartItemData)
	if err != nil {
		return models.CartItem{}, err
	}

	response := models.CartItem{
		Id:        result.Id,
		CartId:    result.CartId,
		ProductId: result.ProductId,
		SizeId:    result.SizeId,
		VariantId: result.VariantId,
		Quantity:  result.Quantity,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	return response, nil
}

func (s *CartItemService) GetCartItemsByCartId(cartId int) ([]models.CartItem, error) {
	result, err := s.repo.GetCartItemsByCartId(cartId)
	if err != nil {
		return []models.CartItem{}, err
	}

	return result, nil
}

func (s *CartItemService) DeleteItem(id int, cartId int) (models.CartItem, error) {
	result, err := s.repo.DeleteItem(id, cartId)
	if err != nil {
		return models.CartItem{}, err
	}

	return result, nil
}