package services

import (
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
	result, err := s.repo.AddItem(newCartItemData)
	if err != nil {
		return models.CartItem{}, err
	}

	response := models.CartItem{
		Id:        result.Id,
		CartId:    result.CartId,
		ProductId: result.ProductId,
		Size:      result.Size,
		Hotice:    result.Hotice,
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

func (s *CartItemService) DeleteItem(id int) (models.CartItem, error) {
	result, err := s.repo.DeleteItem(id)
	if err != nil {
		return models.CartItem{}, err
	}

	return result, nil
}
