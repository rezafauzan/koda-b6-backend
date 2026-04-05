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

	if len(newCartItemData.Size) == 0 {
		return models.CartItem{}, errors.New("Failed to add item! : Size is required !")
	}

	if len(newCartItemData.Hotice) == 0 {
		return models.CartItem{}, errors.New("Failed to add item! : Hot/Ice must be filled !")
	}

	if newCartItemData.Hotice != "hot" && newCartItemData.Hotice != "ice" {
		return models.CartItem{}, errors.New("Failed to add item! : Hotice must be either 'hot' or 'ice' !")
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
