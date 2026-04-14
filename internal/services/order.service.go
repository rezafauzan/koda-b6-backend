package services

import (
	"errors"
	"rezafauzan/koda-b6-golang/internal/dto"
	"rezafauzan/koda-b6-golang/internal/repository"
	"strings"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreatePayment(userID int, req dto.CreatePaymentRequestDTO) (map[string]any, error) {
	if userID <= 0 {
		return nil, errors.New("Invalid user session")
	}
	if req.Fullname == "" || len(req.Fullname) < 4 {
		return nil, errors.New("Checkout failed : Fullname minimum 4 characters")
	}
	if req.Phone == "" || len(req.Phone) < 10 {
		return nil, errors.New("Checkout failed : Phone number minimum 10 digits")
	}
	if req.Email == "" || !strings.Contains(req.Email, "@") {
		return nil, errors.New("Checkout failed : Invalid email")
	}
	if req.Address == "" || len(req.Address) < 10 {
		return nil, errors.New("Checkout failed : Address minimum 10 characters")
	}
	if req.Delivery == "" || len(req.Delivery) < 4 {
		return nil, errors.New("Checkout failed : Delivery method required")
	}
	return s.repo.CheckoutOrder(userID, req)
}
