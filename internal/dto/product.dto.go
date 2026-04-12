package dto

import "time"

type ProductResponseDTO struct {
	Id          int       `json:"id"`
	CategoryId  int       `json:"category_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64       `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateProductRequestDTO struct {
	Name        string  `json:"name" binding:"required,min=4"`
	CategoryId  int     `json:"category_id" binding:"required,gt=0"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"gte=0"`
}
